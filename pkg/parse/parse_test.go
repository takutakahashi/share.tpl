package parse

import (
	"reflect"
	"testing"
)

func TestExecute(t *testing.T) {
	type args struct {
		in   []byte
		data map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				in: []byte("hello @@(name), are you @@(hangly)?"),
				data: map[string]string{
					"name":   "bob",
					"hangly": "HANGLY",
				},
			},
			want:    []byte("hello bob, are you HANGLY?"),
			wantErr: false,
		},
		{
			name: "ng",
			args: args{
				in: []byte("hello @@(name), are you @@(hangly)?"),
				data: map[string]string{
					"name": "bob",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Execute(tt.args.in, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
