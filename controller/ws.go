package controller

import (
	"Chess/dao"
	"Chess/module"
	"Chess/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"sync"
)

var clients = make(map[module.WsKey]*websocket.Conn)

//广播通道
var broadcast = make(chan module.UserMessage, 10000)

var lock sync.RWMutex

var UpGrader = websocket.Upgrader{
	//跨域设置
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func InitChat() {
	//启动广播推送协程
	go pushMessages()
	//启动消息批量处理协程，借助消息队列将消息批量插入数据库
	go insertMessage()
}

func ChatLink(c *gin.Context) {
	username := c.Param("username")
	roomId := c.Param("RID")
	//将id转为int64类型
	uid, _ := strconv.ParseInt(username, 10, 64)
	rid, _ := strconv.ParseInt(roomId, 10, 64)
	//升级
	ws, err := UpGrader.Upgrade(c.Writer, c.Request, nil)
	//将当前连接的客户端放入客户端map（clients）中
	wsKey := module.WsKey{
		RoomId: rid,
		UserId: uid,
	}
	lock.RLock()
	clients[wsKey] = ws
	lock.RUnlock()

	if err != nil {
		fmt.Println(err)
		delete(clients, wsKey) //删除map中的客户端
		return
	}
	defer ws.Close()

	for {
		//读取websocket发来的数据
		_, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			delete(clients, wsKey) //删除map中的客户端
			break
		}

		//创建基础聊天消息模板
		chatMessage := module.Message{}
		chatMessage.MessageData = string(message)                  //聊天消息
		chatMessage.UserId, _ = strconv.ParseInt(username, 10, 64) //用户id
		chatMessage.RoomId, _ = strconv.ParseInt(roomId, 10, 64)
		chatMessage.CreateTime = util.GetLocalDateTime() //消息创建时间
		chatMessage.UpdateTime = util.GetLocalDateTime() //消息更新时间
		chatMessage.Status = 1                           //消息状态

		//查询用户信息
		user := module.SUser{}
		dao.GetDB().Model(&user).Where("id=?", chatMessage.UserId).Find(&user)
		userMessage := module.UserMessage{
			//Message中的数据
			MessageData: chatMessage.MessageData,
			UserId:      chatMessage.UserId,
			RoomId:      chatMessage.RoomId,
			CreateTime:  chatMessage.CreateTime,
			UpdateTime:  chatMessage.UpdateTime,
			Status:      chatMessage.Status,
		}

		//如果消息为空
		if len(chatMessage.MessageData) == 0 {
			//跳过
			continue
		}

		//将消息发送至消息队列
		rmq := util.CreateRabbitMQ("msg", "exc", username)
		byteMsg, err := json.Marshal(userMessage)
		if err == nil {
			rmq.SendMessageRouting(byteMsg)
		}
	}
}

func pushMessages() {
	for {
		//读取通道中的消息
		msg := <-broadcast

		//轮询现有的websocket客户端
		for key, client := range clients {

			//获取RoomId
			gid := key.RoomId

			//匹配客户端，判断该客户端的GroupId是否与该消息的GroupId一致，如果是，则将该消息投递给该客户端
			if msg.RoomId == gid && msg.UserId != 0 && len(msg.MessageData) > 0 {
				//发送消息，含失败重试机制，重试次数:3
				for i := 0; i < 3; i++ {
					//发送消息到消费者客户端
					err := client.WriteJSON(msg)
					//如果发送成功
					if err == nil {
						//结束循环
						break
					}
					//如果到达重试次数，但仍未发送成功
					if i == 2 && err != nil {
						//客户端关闭
						client.Close()
						//删除map中的客户端
						delete(clients, key)
					}
				}
			}
		}
	}
}

//插入消息到数据库
func insertMessage() {
	for {
		//从消息队列中获取消息
		msg, _ := module.GetMsg()

		//将userMessage转成message模板
		chatMessage := module.Message{}
		chatMessage.MessageData = msg.MessageData
		chatMessage.UserId = msg.UserId
		chatMessage.RoomId = msg.RoomId
		chatMessage.CreateTime = util.GetLocalDateTime()
		chatMessage.UpdateTime = util.GetLocalDateTime()
		chatMessage.Status = 1
		tx := dao.GetDB().Begin()

		//将message消息模板插入数据库
		err := dao.GetDB().Create(&chatMessage).Error
		if err != nil {
			//错误-事务回滚
			tx.Rollback()
			continue
		}


		//事务提交
		tx.Commit()
	}
}
