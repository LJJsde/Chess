package client

import "testing"

func TestCreateClient(t *testing.T) {
	tests := []struct {
		name string
	}{{
		//这个函数不需要传参数
		"1",
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateClient()
		})
	}
}
