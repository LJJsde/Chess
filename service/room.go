package service

import (
	"Chess/module"
	"log"
	"sync"
)

var RoomCountAuto = 0
var RoomInfo sync.Map

type Room struct {
	*module.SRoom
}

func NewRoom() int {
	RoomCountAuto++
	nIndex := RoomCountAuto
	r := Room{}
	r.RID = RoomCountAuto
	RoomInfo.Store(nIndex, r)
	return r.RID
}

func CheckNewRoom() module.SRoom {
	var pRoom module.SRoom
	RoomInfo.Range(func(k, v interface{}) bool {
		if v.(*module.SRoom).CheckEmpty() {
			pRoom = v.(*module.SRoom)
			return false
		}
		return true
	})
	return pRoom
}

func (i *Room) StartPlay() {
	go func() {
		log.Println("Game start!")
		//i.Chess = game.NewChess()
		log.Println("Game Over!")
	}()
}

func (i *Room) init() {
	i.PlayerNum = 0
}

func (i *Room) CheckFull() bool {
	return i.PlayerNum >= 2
}

func (i *Room) CheckEmpty() bool {
	return i.PlayerNum < 2
}

func (self *Room) JoinRoom(Player *module.SUser) {
	self.PlayerNum++
	for i := 0; i < 3; i++ {
		if self.PlayerList[i] == nil {
			self.PlayerList[i] = Player
			return
		}
	}
}

func QuickJoin(Player *module.SUser, err error) bool {
	fRoom := CheckNewRoom()
	if fRoom == nil {
		log.Println("There's no empty room,creating new room")
		NewRoom()
	}
	fRoom.JoinRoom(Player)
	if fRoom.CheckEmpty() {
		fRoom.StartPlay()
	}
	if err != nil {
		log.Println(err)
	}
	return true
}
