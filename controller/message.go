package controller

import (
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var tempChat = map[string][]common.Message{}

var messageIdSequence = int64(1)

type ChatResponse struct {
	common.Response
	MessageList []common.Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	userInfo, err := service.GetUserInfo(token)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	toUserId := c.Query("to_user_id")
	content := c.Query("content")
	fromId := userInfo.Id
	toId, _ := strconv.Atoi(toUserId)
	if err := service.MessageAdd(fromId, int64(toId), content); err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 2, StatusMsg: err.Error()})
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 0})
	}
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	userInfo, err := service.GetUserInfo(token)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	toUserId := c.Query("to_user_id")
	toId, _ := strconv.Atoi(toUserId)
	preMsgTime, _ := strconv.Atoi(c.Query("pre_msg_time"))
	fromId := userInfo.Id
	messages := service.MessageList(fromId, int64(toId), int64(preMsgTime))
	c.JSON(http.StatusOK, ChatResponse{Response: common.Response{StatusCode: 0}, MessageList: messages})
}
