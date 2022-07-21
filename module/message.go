package module

import (
	"Chess/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
)

var (
	enteringChannel = make(chan *RUser)
	leavingChannel  = make(chan *RUser)
	messageChannel  = make(chan string, 5)
)

func sendMessageToMQ(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func broadcaster() {
	// 使用结构体map来存储所有用户信息
	users := make(map[*RUser]struct{})
	for {
		select {
		case user := <-enteringChannel:
			// 新用户进入保存到map里面
			users[user] = struct{}{}
		case user := <-leavingChannel:
			// 用户退出后,删除掉这个用户
			delete(users, user)
		case msg := <-messageChannel:
			// 给在线的用户发送消息
			for user := range users {
				user.MessageChannel <- msg
			}
		}
	}
}

//自定义返回内容
func ResponseWithJson(code int, data interface{}, msg error, c *gin.Context) {
	c.JSON(200, &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func GetMsg() (UserMessage, bool) {

	//接收消息队列推送过来的消息mqMsg
	mqMsg, isOk := util.ConsumerReceive()

	//取出消息队列的消息主体，将其解析到models.UserMessage
	userMessage := UserMessage{}
	if isOk {
		err := json.Unmarshal([]byte(mqMsg), &userMessage)
		if err != nil {
			return userMessage, false
		}
		return userMessage, true
	}
	return userMessage, false
}
