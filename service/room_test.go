package service

import "testing"

func TestNewRoom(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRoom(); got != tt.want {
				t.Errorf("NewRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}
