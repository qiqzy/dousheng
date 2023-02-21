package controller

import (
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

type UserLoginResponse struct {
	common.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	common.Response
	User common.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	newUser, err := service.Register(username, password)

	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 0},
			UserId:   newUser.Id,
			Token:    newUser.Token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	user, err := service.Login(username, password)

	if err == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    user.Token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	userinfo, err := service.GetUserInfo(token)
	userinfo.Avatar = "http://" + viper.GetString("video.ip") + ":8080/static/" + "avatar.jpg"
	userinfo.BackgroundImage = "http://" + viper.GetString("video.ip") + ":8080/static/" + "background.jpg"
	userinfo.Signature = "执子之手，将子拖走"
	if err == nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 0},
			User:     *userinfo,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}
}
