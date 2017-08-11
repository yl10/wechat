package client

import (
	"fmt"
	"testing"
)

var QYH *Client

func TestMain(t *testing.T) {
	qyh, err := NewClient("wx0cf1b9bacb973d32", "q0O-LRbvIX4LcVaKMQbJXl87xXYlbbWhJss-5Tn2blOl2PQ5ghbMiHsskz_KtKDQ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(qyh.AccessToken())
	}
	QYH = qyh
	ds, err := QYH.GetDepartmentlist()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ds)
	}

}

func TestDepartment(t *testing.T) {
	id, err := QYH.CreateDepartment("bumen1", "1", "1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(QYH.UpdateDepartment(id, "bumen11", "1", "1"))
	}
	ds, err := QYH.GetDepartmentlist()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ds)
	}
}
