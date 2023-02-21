package controller

import (
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	common.Response
	CommentList []common.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	common.Response
	Comment common.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	userInfo, err := service.GetUserInfo(token)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	actionType := c.Query("action_type")

	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	switch actionType {
	case "1":
		commentText := c.Query("comment_text")
		comment, err := service.CommentAdd(userInfo.Id, int64(videoId), commentText)
		if err != nil {
			c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		c.JSON(http.StatusOK, CommentActionResponse{Response: common.Response{StatusCode: 0},
			Comment: *comment,
		})
	case "2":
		commentId, err := strconv.Atoi(c.Query("comment_id"))
		if err != nil {
			c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		err = service.CommentDelete(userInfo.Id, int64(videoId), int64(commentId))
		if err != nil {
			c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.Response{StatusCode: 0})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	comments, err := service.CommentList(int64(videoId))
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    common.Response{StatusCode: 0},
		CommentList: comments,
	})
}
