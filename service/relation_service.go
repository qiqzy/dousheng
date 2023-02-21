package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/model"
	"strconv"
)

func FollowAction(fromUserId int64, toUserId int64, actionType int32) error {
	if fromUserId == toUserId {
		return nil
	}
	has, err := model.QueryFollowByFollowerIdFollowId(fromUserId, toUserId)
	if err != nil {
		return err
	}
	delta := 0
	switch actionType {
	case 1:
		if has {
			return nil
		}
		err = model.AddFollow(fromUserId, toUserId)
		if err != nil {
			return err
		}
		delta = 1
	case 2:
		if !has {
			return nil
		}
		err = model.DeleteFollow(fromUserId, toUserId)
		if err != nil {
			return err
		}
		delta = -1
	default:
		return errors.New("follower action :" + strconv.Itoa(int(actionType)) + "is not correct")
	}
	if delta == 0 {
		return nil
	}

	user, err := model.QueryUserById(int(fromUserId))
	if err != nil {
		return nil
	}
	user.FollowCount += int64(delta)
	err = model.UpdateUser(*user)
	if err != nil {
		return err
	}

	user, err = model.QueryUserById(int(toUserId))
	if err != nil {
		return nil
	}
	user.FollowerCount += int64(delta)
	err = model.UpdateUser(*user)
	if err != nil {
		return err
	}
	return nil
}

func FollowList(userId int64) ([]common.User, error) {
	var res []common.User
	follows, err := model.QueryFollowByFollowerId(userId)
	if err != nil {
		return res, err
	}
	n := len(follows)
	for i := 0; i < n; i++ {
		follow := follows[i]
		user, err := model.QueryUserById(int(follow.FollowID))
		if err != nil {
			return res, err
		}
		now := common.User{Id: int64(user.ID), Name: user.Username, FollowCount: user.FollowCount, FollowerCount: user.FollowerCount}
		now.IsFollow = false //to modify later
		res = append(res, now)
	}
	return res, nil
}

func FollowerList(userId int64) ([]common.User, error) {
	var res []common.User
	followers, err := model.QueryFollowByFollowId(userId)
	if err != nil {
		return res, err
	}
	n := len(followers)
	for i := 0; i < n; i++ {
		follower := followers[i]
		user, err := model.QueryUserById(int(follower.FollowerID))
		if err != nil {
			return res, err
		}
		now := common.User{Id: int64(user.ID), Name: user.Username, FollowCount: user.FollowCount, FollowerCount: user.FollowerCount}
		now.IsFollow = false //to modify later
		res = append(res, now)
	}
	return res, nil
}

func FriendList(userId int64) ([]common.User, error) {
	var res []common.User
	follows, err := model.QueryFollowByFollowerId(userId)
	if err != nil {
		return res, err
	}
	followers, err := model.QueryFollowByFollowId(userId)
	if err != nil {
		return res, err
	}
	for _, follow := range follows {
		flag := false
		for _, follower := range followers {
			if follow.FollowID == follower.FollowerID {
				flag = true
				break
			}
		}
		if flag {
			user, err := model.QueryUserById(int(follow.FollowID))
			if err != nil {
				return res, err
			}
			now := common.User{Id: int64(user.ID), Name: user.Username, FollowCount: user.FollowCount, FollowerCount: user.FollowerCount}
			now.IsFollow = false //to modify later
			res = append(res, now)
		}
	}
	return res, nil
}
