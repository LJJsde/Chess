package service

import (
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
)

var onlineUsersMap map[string]*SClient
var messageChan chan string

type SConnection struct {
}

type SClient struct {
	c        chan string
	UserName string
	Addr     string
}

var UpGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func CreateChatService() {
	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Println("listen err", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("can't accept client", err)
			return
		}
		go HandleConnect(conn)
	}
}

func Manager() {
	onlineUsersMap = make(map[string]*SClient)
	for {
		msg := <-messageChan
		for _, SClient := range onlineUsersMap {
			SClient.c <- msg
		}
	}
}

func HandleConnect(conn net.Conn) {
	defer conn.Close()
	netAddr := conn.RemoteAddr().String()
	client := &SClient{make(chan string), netAddr, netAddr}
	onlineUsersMap[netAddr] = client
	go writeMsgToClient(client, conn)
	loginMsg := "[" + client.Addr + "]" + client.UserName + "login!\n"
	messageChan <- loginMsg
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				log.Println("user[%s]:%shas been quit\n", client, netAddr, client.UserName)
				return
			}
			if err != nil {
				log.Println("cant read info", err)
				return
			}
			msg := "[" + client.Addr + "]" + client.UserName + ":" + string((buf[:n-1])) + "\n"
			messageChan <- msg
		}
	}()

	for {
	}
}

func writeMsgToClient(client *SClient, conn net.Conn) {
	for msg := range client.c {
		conn.Write([]byte(msg))
	}
}
