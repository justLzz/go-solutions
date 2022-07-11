package message

type Message struct {
	Type string
	Data string
}

const (
	LoginMesType = "loginMes"
	LoginResMesType = "loginResMes"
)

