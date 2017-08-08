package addr

import (
	"reflect"
	"testing"

	"github.com/yl10/wechat/qy/client"
)

func TestCreateDepartment(t *testing.T) {
	type args struct {
		c        *client.Client
		name     string
		parentid string
		order    string
	}
	tests := []struct {
		name    string
		args    args
		wantId  int
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := CreateDepartment(tt.args.c, tt.args.name, tt.args.parentid, tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDepartment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("CreateDepartment() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestUpdateDepartment(t *testing.T) {
	type args struct {
		c        *client.Client
		id       int
		name     string
		parentid string
		order    string
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
			if err := UpdateDepartment(tt.args.c, tt.args.id, tt.args.name, tt.args.parentid, tt.args.order); (err != nil) != tt.wantErr {
				t.Errorf("UpdateDepartment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteDepartment(t *testing.T) {
	type args struct {
		c  *client.Client
		id int
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
			if err := DeleteDepartment(tt.args.c, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteDepartment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetDepartmentlist(t *testing.T) {
	type args struct {
		c  *client.Client
		id []int
	}
	tests := []struct {
		name    string
		args    args
		want    []Department
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDepartmentlist(tt.args.c, tt.args.id...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDepartmentlist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDepartmentlist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	type args struct {
		c    *client.Client
		user User
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
			if err := CreateUser(tt.args.c, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateUserFull(t *testing.T) {
	type args struct {
		c          *client.Client
		userid     string
		name       string
		department []int
		position   string
		mobile     string
		email      string
		weixinid   string
		attrs      []UserAttr
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
			if err := CreateUserFull(tt.args.c, tt.args.userid, tt.args.name, tt.args.department, tt.args.position, tt.args.mobile, tt.args.email, tt.args.weixinid, tt.args.attrs); (err != nil) != tt.wantErr {
				t.Errorf("CreateUserFull() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	type args struct {
		c      *client.Client
		userid string
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
			if err := DeleteUser(tt.args.c, tt.args.userid); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBatchDeleteUser(t *testing.T) {
	type args struct {
		c          *client.Client
		useridlist []string
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
			if err := BatchDeleteUser(tt.args.c, tt.args.useridlist); (err != nil) != tt.wantErr {
				t.Errorf("BatchDeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserinfo(t *testing.T) {
	type args struct {
		c      *client.Client
		userid string
	}
	tests := []struct {
		name    string
		args    args
		want    User
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserinfo(tt.args.c, tt.args.userid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserinfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserinfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserListByDept(t *testing.T) {
	type args struct {
		c          *client.Client
		deptID     string
		details    bool
		fetchChild []bool
	}
	tests := []struct {
		name    string
		args    args
		want    []User
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserListByDept(tt.args.c, tt.args.deptID, tt.args.details, tt.args.fetchChild...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserListByDept() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserListByDept() = %v, want %v", got, tt.want)
			}
		})
	}
}
