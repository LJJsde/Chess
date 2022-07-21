package module

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

//导出位置的结构体
var (
	X int
	Y int
)

//用于自定义返回数据的结构体
type Response struct {
	Code int
	Msg  error
	Data interface{}
}

//房间
type SRoom struct {
	RID        int
	HasFull    bool
	PlayerNum  int
	PlayerList [2]*SUser
	Chess      bool
}

//游客
var (
	ErrDuplicateEmail = errors.New("duplicate email")
	AnonymousUser     = &SUser{}
)

//用户
type SUser struct {
	gorm.Model
	UserID   string `gorm:"varchar(10);not null;unique"`
	UserName string `gorm:"varchar(20);not null"`
	Password string `gorm:"size:10;not null"`
	Email    string `gorm:"varchar(20);not null;unique"`
}

type RUser struct {
	CurrentInRoom  int
	EnterChatTime  time.Time
	MessageChannel chan string
}

//消息的传输版本
type UserMessage struct {
	Id          int64  `gorm:"column:id;primary_key;auto_increment" json:"id"`
	RoomId      int64  `gorm:"column:room_id" json:"roomId"`
	UserId      int64  `gorm:"column:user_id" json:"userId"`
	MessageData string `gorm:"column:message_data" json:"messageData"`
	CreateTime  string `gorm:"column:create_time" json:"createTime"`
	UpdateTime  string `gorm:"column:update_time" json:"updateTime"`
	Status      int64  `gorm:"column:status;" json:"status"`
}

//房间内部消息
type Message struct {
	Id          int64  `gorm:"column:id;primary_key;auto_increment" json:"id"`
	RoomId      int64  `gorm:"column:room_id" json:"roomId"`
	UserId      int64  `gorm:"column:user_id" json:"userId"`
	MessageData string `gorm:"column:message_data" json:"messageData"`
	CreateTime  string `gorm:"column:create_time" json:"createTime"`
	UpdateTime  string `gorm:"column:update_time" json:"updateTime"`
	Status      int64  `gorm:"column:status;" json:"status"`
}
type WsKey struct {
	RoomId int64
	UserId int64
}
