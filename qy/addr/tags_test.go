package addr

import (
	"reflect"
	"testing"

	"github.com/yl10/wechat/qy/client"
)

var wx *client.Client

func init() {
	wx, _ = client.DefaultTestClient()
}
func TestCreateTag(t *testing.T) {
	type args struct {
		c    *client.Client
		name string
		id   []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "创建标签",
			args: args{
				c:    wx,
				name: "测试分组",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateTag(tt.args.c, tt.args.name, tt.args.id...)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if got != tt.want {
			// 	t.Errorf("CreateTag() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestUpdateTag(t *testing.T) {
	type args struct {
		c    *client.Client
		name string
		id   int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateTag(tt.args.c, tt.args.name, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteTag(t *testing.T) {
	type args struct {
		c     *client.Client
		tagid int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteTag(tt.args.c, tt.args.tagid); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUsersByTag(t *testing.T) {
	type args struct {
		c     *client.Client
		tagid int
	}
	tests := []struct {
		name          string
		args          args
		wantUserlist  map[string]string
		wantPartylist []int
		wantErr       bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUserlist, gotPartylist, err := GetUsersByTag(tt.args.c, tt.args.tagid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsersByTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUserlist, tt.wantUserlist) {
				t.Errorf("GetUsersByTag() gotUserlist = %v, want %v", gotUserlist, tt.wantUserlist)
			}
			if !reflect.DeepEqual(gotPartylist, tt.wantPartylist) {
				t.Errorf("GetUsersByTag() gotPartylist = %v, want %v", gotPartylist, tt.wantPartylist)
			}
		})
	}
}

func TestAddUserToTag(t *testing.T) {
	type args struct {
		c         *client.Client
		tagid     int
		userlist  []string
		partylist []int
	}
	tests := []struct {
		name             string
		args             args
		wantInvalidlist  []string
		wantInvalidparty []int
		wantErr          bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInvalidlist, gotInvalidparty, err := AddUserToTag(tt.args.c, tt.args.tagid, tt.args.userlist, tt.args.partylist)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddUserToTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInvalidlist, tt.wantInvalidlist) {
				t.Errorf("AddUserToTag() gotInvalidlist = %v, want %v", gotInvalidlist, tt.wantInvalidlist)
			}
			if !reflect.DeepEqual(gotInvalidparty, tt.wantInvalidparty) {
				t.Errorf("AddUserToTag() gotInvalidparty = %v, want %v", gotInvalidparty, tt.wantInvalidparty)
			}
		})
	}
}

func TestDeleteUserFromTag(t *testing.T) {
	type args struct {
		c         *client.Client
		tagid     int
		userlist  []string
		partylist []int
	}
	tests := []struct {
		name             string
		args             args
		wantInvalidlist  []string
		wantInvalidparty []int
		wantErr          bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInvalidlist, gotInvalidparty, err := DeleteUserFromTag(tt.args.c, tt.args.tagid, tt.args.userlist, tt.args.partylist)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteUserFromTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInvalidlist, tt.wantInvalidlist) {
				t.Errorf("DeleteUserFromTag() gotInvalidlist = %v, want %v", gotInvalidlist, tt.wantInvalidlist)
			}
			if !reflect.DeepEqual(gotInvalidparty, tt.wantInvalidparty) {
				t.Errorf("DeleteUserFromTag() gotInvalidparty = %v, want %v", gotInvalidparty, tt.wantInvalidparty)
			}
		})
	}
}

func Test_patchuserbytag(t *testing.T) {
	type args struct {
		c         *client.Client
		action    Action
		tagid     int
		userlist  []string
		partylist []int
	}
	tests := []struct {
		name             string
		args             args
		wantInvalidlist  []string
		wantInvalidparty []int
		wantErr          bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInvalidlist, gotInvalidparty, err := patchuserbytag(tt.args.c, tt.args.action, tt.args.tagid, tt.args.userlist, tt.args.partylist)
			if (err != nil) != tt.wantErr {
				t.Errorf("patchuserbytag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInvalidlist, tt.wantInvalidlist) {
				t.Errorf("patchuserbytag() gotInvalidlist = %v, want %v", gotInvalidlist, tt.wantInvalidlist)
			}
			if !reflect.DeepEqual(gotInvalidparty, tt.wantInvalidparty) {
				t.Errorf("patchuserbytag() gotInvalidparty = %v, want %v", gotInvalidparty, tt.wantInvalidparty)
			}
		})
	}
}

func TestGetTagList(t *testing.T) {
	type args struct {
		c *client.Client
	}
	tests := []struct {
		name    string
		args    args
		want    map[int]string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "测试获取标签",
			args: args{
				c: wx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTagList(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTagList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("GetTagList() = %v, want %v", got, tt.want)
			// }
			t.Logf("获取到的标签为：\r\n%v\r\n", got)
		})
	}
}
