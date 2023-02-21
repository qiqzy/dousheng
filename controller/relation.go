package controller

import (
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	common.Response
	UserList []common.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserId, err := strconv.Atoi(c.Query("to_user_id"))
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
	err = service.FollowAction(userInfo.Id, int64(toUserId), int32(actionType))
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.Response{StatusCode: 0})

}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	userInfo, err := service.GetUserInfo(token)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	users, err := service.FollowList(userInfo.Id)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, UserListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	userInfo, err := service.GetUserInfo(token)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	users, err := service.FollowerList(userInfo.Id)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, UserListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	token := c.Query("token")
	userInfo, err := service.GetUserInfo(token)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	users, err := service.FriendList(userInfo.Id)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, UserListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}
