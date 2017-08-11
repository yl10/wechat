package wxencrypt

import "testing"

func Test_getRandomStr(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{name: "随机字符串测试"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%s", getRandomStr())
		})
	}
}
