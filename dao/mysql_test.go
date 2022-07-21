package dao

import (
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestConnDB(t *testing.T) {
	tests := []struct {
		name string
		want *gorm.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConnDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateNewUser(t *testing.T) {
	type args struct {
		email          string
		hashedpassword []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{email: "9999999999@qq.com", hashedpassword: []byte{0x12, 0x13, 0x14, 0x12, 0x13}}},
		{"2", args{email: "", hashedpassword: []byte{}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateNewUser(tt.args.email, tt.args.hashedpassword)
		})
	}
}
