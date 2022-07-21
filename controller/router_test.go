package controller

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

func TestInitRouter(t *testing.T) {
	tests := []struct {
		name string
		want *gin.Engine
	}{
		{ "1",}
	},
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Register(tt.args.ctx)
		})
	}
}

func TestLogin(t *testing.T) {
	type args struct {
		ctx gin.Context
	}

	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{ctx: ("writermem":"responseWriter"
	"Request":"*http.Request"
	"Writer":"ResponseWriter"
	"Params":"Params"
	"handlers":"HandlersChain"
	"index":"int8"
	"fullPath":"string"
	"engine":"*Engine"
	"params":"*Params"
	"skippedNodes":"*[]skippedNode"
	"mu":"sync.RWMutex"
	"Keys":"map[string]interface{}"
	"Errors":"errorMsgs"
    "Accepted":"[]string"
    "queryCache":"url.Values"
    "formCache":"url.Values"
    "sameSite":"http.SameSite")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Login(tt.args.ctx)
		})
	}
}
