package foo

import (
	"context"
	"errors"
	"fmt"

	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"
	"zero-demo/user-api/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UserCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCreateLogic) UserCreate(req *types.UserCreateReq) (resp *types.UserCreateResp, err error) {
	// 开启事务
	if err := l.svcCtx.UserModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		//1 ADD User
		user := &model.User{}
		user.Mobile = req.Mobile
		user.Nickname = req.NickName

		dbResult, err := l.svcCtx.UserModel.TransInsert(ctx, session, user)
		if err != nil {
			return err
		}
		
		userId, _ := dbResult.LastInsertId()

		// 2 ADD userdata
		userData := &model.UserData{}
		userData.UserId = userId
		userData.Data = "xxxxxxxxxxxxxx"

		if _, err := l.svcCtx.UserDataModel.TransInsert(ctx, session, userData); err != nil {
			fmt.Println("[+++++++++++]err: ", err.Error())
			return err
		}

		return nil

	}); err != nil {
		return nil, errors.New("创建用户失败")
	}

	return &types.UserCreateResp{
		Flag: true,
	}, nil
}
