package module

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

func TestResponseWithJson(t *testing.T) {
	type args struct {
		code int
		data interface{}
		msg  error
		c    *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ResponseWithJson(tt.args.code, tt.args.data, tt.args.msg, tt.args.c)
		})
	}
}

func TestGetMsg(t *testing.T) {
	tests := []struct {
		name  string
		want  UserMessage
		want1 bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetMsg()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMsg() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetMsg() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
