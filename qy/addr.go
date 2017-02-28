package qy

import "github.com/yl10/wechat/qy/addr"

//GetUserInfoByUserID 通过userID获取用户信息
func (wx *Wx) GetUserInfoByUserID(userID string) (addr.User, error) {
	return addr.GetUserinfo(wx.client, userID)
}
