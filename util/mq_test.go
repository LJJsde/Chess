package util

import (
	"reflect"
	"testing"
)

func TestCreateRabbitMQ(t *testing.T) {
	type args struct {
		queuename string
		exchange  string
		key       string
	}
	tests := []struct {
		name string
		args args
		want *RabbitMQ
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateRabbitMQ(tt.args.queuename, tt.args.exchange, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRabbitMQ() = %v, want %v", got, tt.want)
			}
		})
	}
}
