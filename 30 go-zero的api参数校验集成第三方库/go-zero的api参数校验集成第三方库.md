## 0 安装validate
```go
go get github.com/go-playground/validator
```

## 1 修改API的请求结构体，增加标签
```go
// type literal
type (
	UserInforReq {
		UserId int64 `json:"userId" validate:"gte:30 lte:40"`
	}
	UserInforResp {
		UserId   int64  `json:"userId"`
		NickName string `json:"nickName"`
	}
)
```
> 参考语法 go get github.com/go-playground/validator

## 2 重新生成api，会在type中看到

## 3 在handler中加载validate处理
```go
func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInforReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
        // 增加validate处理------------------------------
        if err := validate.New().StructCtx(r.context,req);err != nil {
            httpx.Error(w, err)
			return
        }
        //----------------------------------------------

		l := foo.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

```
