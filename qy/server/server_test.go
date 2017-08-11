package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTP(t *testing.T) {
	s, err := NewServer("GMccroGSfrcHG3", "vLjKCNVAxyavF6FkguEJW2TbJ84M0ky8IBw28HkNSne", "wx59ab475b9a833d82")
	fmt.Println(err)

	s.RegisterHandler(ActionUpdateTag, func(w http.ResponseWriter, msg RequestMsg) error {
		fmt.Println("更新tag")
		return nil
	})

	s.RegisterHandler(ActionEnterAgent, func(w http.ResponseWriter, msg RequestMsg) error {
		fmt.Println("jinru jinyong")
		return nil
	})

	hs := httptest.NewServer(s)
	t.Log(hs.URL)
}
