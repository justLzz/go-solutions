package message

type LoginMes struct {
	UserAccount string
	UserPwd string
	UserName string
}

type LoginResMes struct {
	Code int
	Error string
}
