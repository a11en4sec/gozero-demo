package foo

import (
	"context"

	"github.com/pkg/errors"

	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"

	"zero-demo/user-api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInforReq) (resp *types.UserInforResp, err error) {
	// todo: add your logic here and delete this line

	if err := l.testOne(); err != nil {
		logx.Errorf("err: %+v\n", err)
	}
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)

	if err != nil && err != model.ErrNotFound {
		return nil, errors.New("查询数据库失败")
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	return &types.UserInforResp{
		UserId:   user.Id,
		NickName: user.Nickname,
	}, nil
}

func (l *UserInfoLogic) testOne() error {
	return l.testTwo()
}

func (l *UserInfoLogic) testTwo() error {
	return l.testThree()
}

func (l *UserInfoLogic) testThree() error {
	return errors.Wrap(errors.New("故意报错"), "测试堆栈打印")
}
