package global

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestLoadSetting(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    SnipConfig
		wantErr bool
	}{
		{
			name: "ok",
			args: args{filepath: "../../misc/config_test.yaml"},
			want: SnipConfig{
				Setting: Setting{
					BaseDir: "misc",
				},
				Repositories: []Repository{
					{
						Name: "snippets",
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
				b, _ := yaml.Marshal(tt.want)
				t.Logf("%s", string(b))
				t.Errorf("LoadSetting() = %v, want %v", got, tt.want)
			}
		})
	}
}
