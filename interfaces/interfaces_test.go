package interfaces

import (
	"bytes"
	"io"
	"testing"
)

func TestCopy(t *testing.T) {
	type args struct {
		in io.ReadSeeker
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}{
		{
			name: "base-case",
			args: args{
				in: bytes.NewReader([]byte("hello")),
			},
			wantOut: "hellohello",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			if err := Copy(tt.args.in, out); (err != nil) != tt.wantErr {
				return
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("Copy() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
