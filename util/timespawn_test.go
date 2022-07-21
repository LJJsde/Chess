package util

import "testing"

func TestGetLocalDateTime(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLocalDateTime(); got != tt.want {
				t.Errorf("GetLocalDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
