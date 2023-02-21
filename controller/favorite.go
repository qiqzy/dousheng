package controller

import (
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	actionType, err := strconv.Atoi(c.Query("action_type"))
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	userInfo, err := service.GetUserInfo(token)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	err = service.FavoriteAction(userInfo.Id, int64(videoId), int32(actionType))
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, common.Response{StatusCode: 0})
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")

	userInfo, err := service.GetUserInfo(token)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	videos, err := service.FavoriteList(userInfo.Id)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
