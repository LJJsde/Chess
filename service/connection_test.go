package service

import (
	"net"
	"testing"
)

func TestCreateChatService(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateChatService()
		})
	}
}

func TestManager(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Manager()
		})
	}
}

func TestHandleConnect(t *testing.T) {
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandleConnect(tt.args.conn)
		})
	}
}

func Test_writeMsgToClient(t *testing.T) {
	type args struct {
		client *SClient
		conn   net.Conn
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writeMsgToClient(tt.args.client, tt.args.conn)
		})
	}
}
