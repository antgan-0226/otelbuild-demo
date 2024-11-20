package service

import (
	"context"
	"otelbuild-demo/model"
)

type User struct {
	Id       int64
	Nickname string
}

func CreateUser(ctx context.Context, id int64, nickname string) (u *model.UserModel, err error) {
	user := &model.UserModel{
		Id:       id,
		Nickname: nickname,
	}
	return user, nil
}
