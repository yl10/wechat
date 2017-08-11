package jssdk

import "testing"

func Test_Jssha1(t *testing.T) {
	type args struct {
		noncestr    string
		jsapiTicket string
		timestamp   int64
		url         string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "测试签名计算",
			args: args{
				noncestr:    "Wm3WZYTPz0wzccnW",
				jsapiTicket: "sM4AOVdWfPE4DxkXGEs8VMCPGGVi4C3VM0P37wVUCFvkVAy_90u5h9nbSlYy3-Sl-HhTdfl2fzFy1AOcHKP7qg",
				timestamp:   1414587457,
				url:         "http://mp.weixin.qq.com/",
			},
			want: "a1a0141122148a10c58bbd00ef52231dc4dda87a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Jssha1(tt.args.noncestr, tt.args.jsapiTicket, tt.args.timestamp, tt.args.url); got != tt.want {
				t.Errorf("jssha1() = %v, want %v", got, tt.want)
			}
		})
	}
}
