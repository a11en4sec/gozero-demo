// Code generated by goctl. DO NOT EDIT.
package types

type UserInforReq struct {
	UserId int64 `json:"userId"`
}

type UserInforResp struct {
	UserId   int64  `json:"userId"`
	NickName string `json:"nickName"`
}
