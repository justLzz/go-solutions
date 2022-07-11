package login

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"network/tcp/instant_messaging/simple_chatroom/common/message"
)

func Login(account string, password string) (err error)  {

	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil{
		fmt.Println("链接失败：", err)
		return
	}

	defer conn.Close()

	var  mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserAccount = account
	loginMes.UserPwd = password

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("loginMes序列化错误：", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes序列化错误：", err)
		return
	}

	dataLen := uint32(len(data))
	lenByte := make([]byte, 8064)
	binary.BigEndian.PutUint32(lenByte, dataLen)
	_, err = conn.Write(lenByte)
	if err != nil {
		fmt.Println("发送数据长度失败：", err)
		return
	}

	fmt.Println("发送数据长度成功：", len(data))
	return
}
