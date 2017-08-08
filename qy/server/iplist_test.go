package server

import (
	"testing"

	"github.com/yl10/wechat/qy/client"
)

func TestGetCallbackIP(t *testing.T) {
	type args struct {
		c *client.Client
	}
	wx, _ := client.DefaultTestClient()
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "测试获取ip",
			args:    args{wx},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCallbackIP(tt.args.c)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetCallbackIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("GetCallbackIP() = %v, want %v", got, tt.want)
			// }

			t.Logf("获取到的IP为：%v", got)
		})
	}
}
