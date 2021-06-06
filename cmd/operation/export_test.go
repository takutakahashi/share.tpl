package operation

import (
	"reflect"
	"testing"
)

func Test_exportDir(t *testing.T) {
	type args struct {
		opt ExportOpt
	}
	tests := []struct {
		name    string
		args    args
		want    ExportOut
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				opt: ExportOpt{
					Path:          "src/dirtest",
					OutputDirPath: "dist/dirtest",
					Data:          map[string]string{"name": "bob"},
				},
			},
			want: ExportOut{
				Files: map[string][]byte{
					"src/dirtest/src/test.py": []byte(`print("bob")`),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := exportDir(tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("exportDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exportDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
