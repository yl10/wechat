package model

type User struct {
	Userid     string
	Name       string
	Department []int
	Position   string
	Mobile     string
	Email      string
	Weixinid   string

	Avatar string
	Status int

	Extattr struct {
		Attrs []UserAttr
	}
}

type UserAttr struct {
	Name  string
	value string
}
