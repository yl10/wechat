package server

import (
	"net/http"
)

type QyApp struct {
	token  string
	aeskey string
}

func (app *QyApp) ServerHTTP(w *http.ResponseWriter, r *http.Request) {


if r.Method="Get"
}
//检查签名
func checkSignature(t string, w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()
	var signature string = r.FormValue("msg_signature")
	var timestamp string = r.FormValue("timestamp")
	var nonce string = r.FormValue("nonce")
	strs := sort.StringSlice{t, timestamp, nonce}
	sort.Strings(strs)
	var str string
	for _, s := range strs {
		str += s
	}
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil)) == signature
}
