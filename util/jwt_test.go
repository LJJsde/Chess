package util

import (
	"reflect"
	"testing"
)

func TestGeterateToken(t *testing.T) {
	type args struct {
		id     uint
		mobile string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GeterateToken(tt.args.id, tt.args.mobile)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeterateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GeterateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    *Claims
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
