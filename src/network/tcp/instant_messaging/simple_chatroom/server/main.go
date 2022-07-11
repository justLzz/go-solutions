package main

import (
	"fmt"
	"net"

)

func main()  {
	fmt.Println("监听8889")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("监听错误: ", err)
		return
	}

	for  {
		fmt.Println("等待链接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("链接错误: ", err)
			return
		}
		go process(conn)

	}
}

func process(conn net.Conn)  {
	defer conn.Close()
	for  {
		buf := make([]byte, 4)
		fmt.Println("等待读取客户端数据...")
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("数据读取错误：", err)
			return
		}
		fmt.Println("读到数据：", buf)
		break

	}
}


