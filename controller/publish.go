package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"path/filepath"
	"strings"
)

type VideoListResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")

	userInfo, err := service.GetUserInfo(token)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	videoName := fmt.Sprintf("%d_%s", userInfo.Id, filename)
	saveFile := filepath.Join("./public/", videoName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	ip := viper.GetString("video.ip")
	videoPath := viper.GetString("video.path")
	coverName := strings.Split(videoName, ".")[0] + ".jpg"
	service.Generate(videoPath, videoName, coverName)
	err = service.Publish(userInfo.Id, title, "http://"+ip+":8080/static/"+videoName, "http://"+ip+":8080/static/"+coverName)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, common.Response{
		StatusCode: 0,
		StatusMsg:  videoName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	userInfo, err := service.GetUserInfo(token)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	videos, err := service.List(userInfo.Id)
	fmt.Println(len(videos), videos[0], err)
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
