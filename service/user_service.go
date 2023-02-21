package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/spf13/viper"
)

type UserIdToken struct {
	Id    int64
	Token string
}

func Login(username string, password string) (*UserIdToken, error) {
	user, err := model.QueryUserByName(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user doesn't exist")
	}
	md5Password := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	if md5Password != user.Password {
		return nil, errors.New("password not correct")
	}
	userinfo := common.User{
		Id:            int64(user.ID),
		Name:          user.Username,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      true,
	}
	token, err := Encode(viper.GetString("jwt.secret"), userinfo)
	if err != nil {
		return nil, err
	}
	return &UserIdToken{Id: int64(user.ID), Token: token}, nil
}

func Register(username string, password string) (*UserIdToken, error) {
	user, err := model.QueryUserByName(username)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("user already exist")
	}
	user = &model.User{Username: username, Password: password, Name: username}
	user.Password = fmt.Sprintf("%x", md5.Sum([]byte(user.Password)))
	model.AddUser(*user)
	return Login(username, password)
}

func GetUserInfo(token string) (*common.User, error) {
	return Decode(viper.GetString("jwt.secret"), token)
}
