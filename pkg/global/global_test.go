package global

import (
	"reflect"
	"testing"
)

func TestLoadSetting(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    Setting
		wantErr bool
	}{
		{
			name: "ok",
			args: args{filepath: "../../misc/config_test.yaml"},
			want: Setting{
				Repositories: []Repository{
					{
						Name: "takutaka",
						Type: "git",
						URI:  "ssh://git@github.com:takutakahashi/snippets.git"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadSetting(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadSetting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadSetting() = %v, want %v", got, tt.want)
			}
		})
	}
}
