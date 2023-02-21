package controller

import (
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	videoList, err := service.Feed(time.Now().Unix())
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})
}
