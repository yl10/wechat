package agent

import "testing"

func TestAgent_GetInfoFromTencent(t *testing.T) {
	agent, _ := NewAgent("1000002", "wx59ab475b9a833d82", "xTJKfm8hcUYbHlSzCBdR8NgtLISTmksNDvOVaDiOUis", nil, nil)
	tests := []struct {
		name    string
		a       Agent
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "测试获取应用接口",
			a:    *agent,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.GetInfoFromTencent(); (err != nil) != tt.wantErr {
				t.Errorf("Agent.GetInfoFromTencent() error = %v, wantErr %v", err, tt.wantErr)
			}
			//tt.a.Name = "dddddddd"
			t.Log("partys:")
			t.Log("部门：", tt.a.Partys)
			t.Log("biaoqian:", tt.a.Tags)
			t.Log("reny:", tt.a.Users)
		})
	}
}
