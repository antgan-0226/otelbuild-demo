package biz

import (
	"context"
	"fmt"
	"github.com/antgan-0226/kotelbuild/pkg/api"
	"otelbuild-demo/model"
)

func createUserEnterHook(call api.CallContext, ctx context.Context, id int64, nickname string) {
	fmt.Println(fmt.Sprintf("[service hook]create user enter hook, id :%d, name:%s", id, nickname))

}

func createUserExitHook(call api.CallContext, user *model.UserModel, err error) {
	oldUserId := call.GetParam(1)
	oldNickname := call.GetParam(2)
	fmt.Println(fmt.Sprintf("[service hook]create user exit hook, id :%d, name:%s", oldUserId, oldNickname))
	//赋予新值
	user.Id = 2
	user.Nickname = "ant"
	fmt.Println(fmt.Sprintf("[service hook]create user exit hook, id :%d, name:%s", user.Id, user.Nickname))
}
