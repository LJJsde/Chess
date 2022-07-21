package util

import (
	"testing"
)

func TestFileToByte(t *testing.T) {
	type args struct {
		inPath  string
		outPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"1", args{inPath: "F:", outPath: "F:"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FileToByte(tt.args.inPath, tt.args.outPath); (err != nil) != tt.wantErr {
				t.Errorf("FileToByte() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
