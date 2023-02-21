package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/model"
)

func FavoriteAction(userId int64, videoId int64, actionType int32) error {
	delta := 0
	has, err := model.QueryFavoriteByUserIdVideoId(userId, videoId)
	if err != nil {
		return err
	}
	switch actionType {
	case 1:
		if has {
			return nil
		}
		err := model.AddFavorite(userId, videoId)
		if err != nil {
			return err
		}
		delta = 1
	case 2:
		if !has {
			return nil
		}
		err := model.DeleteFavorite(userId, videoId)
		if err != nil {
			return err
		}
		delta = -1
	default:
		return errors.New("actionType is invalid")
	}
	video, err := model.QueryVideoById(videoId)
	if err != nil {
		return err
	}
	video.FavoriteCount += int64(delta)
	err = model.UpdateVideo(*video)
	return err
}

func FavoriteList(userId int64) ([]common.Video, error) {
	var res []common.Video
	favorites, err := model.QueryFavoriteByuserId(userId)
	if err != nil {
		return res, err
	}
	n := len(favorites)
	for i := 0; i < n; i++ {
		now, err := model.QueryVideoById(favorites[i].VideoID)
		if err != nil {
			return res, err
		}
		aVideo := common.Video{Id: int64(now.ID), PlayUrl: now.PlayUrl, CoverUrl: now.CoverUrl, FavoriteCount: now.FavoriteCount, CommentCount: now.CommentCount}
		user, err := model.QueryUserById(int(now.AuthorID))
		if err != nil {
			return nil, err
		}
		if user == nil {
			return nil, errors.New("The video can't find the author")
		}
		aVideo.Author = common.User{Id: int64(user.ID), Name: user.Name, FollowCount: user.FollowCount, FollowerCount: user.FollowerCount}
		aVideo.Author.IsFollow = true // to modify later
		aVideo.IsFavorite = true      // to modify later
		res = append(res, aVideo)
	}
	return res, nil
}
