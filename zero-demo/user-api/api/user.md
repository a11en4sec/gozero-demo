
### 1. "获取用户信息"

1. 路由定义

- Url: /user/info
- Method: POST
- Request: `UserInforReq`
- Response: `UserInforResp`

2. 请求定义


```golang
type UserInforReq struct {
	UserId int64 `json:"userId"`
}
```


3. 返回定义


```golang
type UserInforResp struct {
	UserId int64 `json:"userId"`
	NickName string `json:"nickName"`
}
```
  

