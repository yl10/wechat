package model

const (
	API_URL_department_create = "https://qyapi.weixin.qq.com/cgi-bin/department/create?access_token=%s"
)

type Department struct {
	Id       int
	Name     string
	Parentid int
	Order    int
	Child    []Department
}
