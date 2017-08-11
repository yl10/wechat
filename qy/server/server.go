package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/yl10/kit/safemap"
	"github.com/yl10/wechat/qy/server/wxencrypt"
)

//HandlerFunc 对消息进行处理的函数
type HandlerFunc func(w http.ResponseWriter, req RequestMsg) error

//Server 服务
type Server struct {
	token          string //token
	encodingAESKey string //消息秘钥
	corpID         string
	encryptHelper  *wxencrypt.WXBizMsgCrypt
	handlerFuncMap *safemap.SafeMap
	path           string //提供服务的urlpath
}

//ServerGroup 服务组
type ServerGroup struct {
	serverList *safemap.SafeMap
	port       int
}

//NewServerGroup 创建新的服务组
func NewServerGroup(port int) *ServerGroup {
	return &ServerGroup{
		serverList: safemap.NewSafeMap(),
		port:       port,
	}
}

//AddServer 添加服务
func (sg ServerGroup) AddServer(s Server) error {
	if s.path == "" {
		return fmt.Errorf("服务没有设置path，不能添加。")
	}
	sg.serverList.Set(s.path, &s)
	return nil
}

//DelServer 服务组中删除对应路径的服务
func (sg ServerGroup) DelServer(path string) {
	sg.serverList.Delete(path)
}

//RegisterHandler 注册处理函数
func (s Server) RegisterHandler(evetype HanderAction, fn HandlerFunc) {

	if s.handlerFuncMap == nil {
		s.handlerFuncMap = safemap.NewSafeMap()
	}
	s.handlerFuncMap.Set(evetype, fn)
}

//NewServer 返回一个新的服务
func NewServer(token, encodingAESKey, corpID string, path ...string) (*Server, error) {
	biz, err := wxencrypt.NewWxBIzMsgCrypt(token, encodingAESKey, corpID)
	if err != nil {
		return nil, err
	}
	urlpath := ""
	if len(path) > 0 {
		urlpath = path[0]
	}
	return &Server{
		token:          token,
		encodingAESKey: encodingAESKey,
		encryptHelper:  biz,
		path:           urlpath,
		handlerFuncMap: safemap.NewSafeMap(),
	}, nil

}

//HandleRequestData 处理请求中的body数据
func (s Server) handleRequestData(w http.ResponseWriter, xmldata []byte) error {

	action, _, result, err := PareXMLMsg(xmldata, false)
	if err != nil {
		return err
	}
	//检查处理方法是否注册过
	var fn HandlerFunc
	var ok bool
	if s.handlerFuncMap.Check(action) {
		fn, ok = s.handlerFuncMap.Get(action).(HandlerFunc)
		if ok {
			switch v := result.(type) {

			default:
				return fn(w, v)
			}

		}
		return fmt.Errorf("内部异常，处理函数不是所期待的 ")
	}
	return fmt.Errorf("接受到没有注册过的请求:%s", action)

}

//VerifyURL 验证URL
func (s Server) VerifyURL(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	msgSignature := r.FormValue("msg_signature")
	timestamp := r.FormValue("timestamp")
	echostr := r.FormValue("echostr")
	nonce := r.FormValue("nonce")

	result, err := s.encryptHelper.VerifyURL(msgSignature, timestamp, nonce, echostr)
	if err == nil {
		w.Write([]byte(result))
	}
}

//HandleRequest 对http请求进行处理
func (s *Server) HandleRequest(w http.ResponseWriter, r *http.Request) error {
	//如果是get方法，进行签名检查
	if r.Method == http.MethodGet {
		s.VerifyURL(w, r)
		return nil
	}
	/*如果是post方法
	先解析出明文消息
	*/
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {

			return err
		}

		msgSignature := r.FormValue("msg_signature")
		timestamp := r.FormValue("timestamp")
		nonce := r.FormValue("nonce")
		defer r.Body.Close()
		data, _ := ioutil.ReadAll(r.Body)

		msgdata, err := s.encryptHelper.DecryptMsg(msgSignature, timestamp, nonce, data)
		if err != nil {
			return err
		}
		//把消息数据送过去处理
		return s.handleRequestData(w, msgdata)
	}
	return nil
}

//ServeHTTP 实现http.handler接口
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := s.HandleRequest(w, r)
	if err != nil {
		log.Printf("处理请求出现错误:%v", err)
	}
}

//ListenAndServe 单个服务启动监听
func (s *Server) ListenAndServe(port int, path ...string) {
	inpath := "/"
	switch {
	case len(path) > 0:
		inpath = path[0]
	case s.path != "":
		inpath = s.path
	}
	http.HandleFunc(inpath, func(w http.ResponseWriter, r *http.Request) {
		s.ServeHTTP(w, r)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), s)

	if err != nil {
		log.Fatalf("监听%d端口失败：%v", port, err)
	}

}

//ListenAndServe 服务组启动监听
func (sg ServerGroup) ListenAndServe(port int) {

	//注册路由
	slist := sg.serverList.Items()
	for _, s := range slist {
		if v, ok := s.(Server); ok {
			http.HandleFunc(v.path, func(w http.ResponseWriter, r *http.Request) {
				v.ServeHTTP(w, r)
			})
		}

		err := http.ListenAndServe(fmt.Sprintf(":%d", port), sg)
		if err != nil {
			log.Fatalf("监听%d端口失败：%v", port, err)
		}
	}
}
func (sg ServerGroup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//获取路由
	inpath := r.RequestURI

	if sg.serverList.Check(inpath) {
		s := sg.serverList.Get(inpath)
		if v, ok := s.(Server); ok {
			v.ServeHTTP(w, r)
		}

	}
	log.Printf("请求的路径没有匹配的服务.url:%s\r\n", inpath)

}
