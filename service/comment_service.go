package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/model"
	"time"
)

func CommentAdd(userId int64, videoId int64, content string) (*common.Comment, error) {
	comment := &model.Comment{UserID: userId, VideoID: videoId, Content: content}
	err := model.AddComment(comment)
	if err != nil {
		return nil, err
	}
	res := &common.Comment{Id: int64(comment.ID), Content: content, CreateDate: comment.CreatedAt.Format(time.Layout)}
	user, err := model.QueryUserById(int(userId))
	if err != nil {
		return nil, err
	}
	res.User = common.User{Id: int64(user.ID), Name: user.Username, FollowCount: user.FollowCount, FollowerCount: user.FollowerCount}
	res.User.IsFollow = false //to modify later

	video, err := model.QueryVideoById(videoId)
	if err != nil {
		return nil, err
	}
	video.CommentCount += 1
	err = model.UpdateVideo(*video)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func CommentDelete(userId int64, videoId int64, commentId int64) error {
	comment, err := model.QueryCommentById(commentId)
	if err != nil {
		return err
	}
	if comment == nil {
		return errors.New("no such comment")
	}
	if comment.UserID != userId || comment.VideoID != videoId {
		return errors.New("comment_id is not you's, don't have the permission")
	}
	err = model.DeleteCommentById(commentId)
	if err != nil {
		return err
	}
	video, err := model.QueryVideoById(videoId)
	if err != nil {
		return err
	}
	video.CommentCount -= 1
	err = model.UpdateVideo(*video)
	if err != nil {
		return err
	}
	return err
}

func CommentList(videoId int64) ([]common.Comment, error) {
	var res []common.Comment
	comments, err := model.QueryCommentByVideoId(videoId)
	if err != nil {
		return res, err
	}
	n := len(comments)
	for i := 0; i < n; i++ {
		comment := comments[i]
		now := common.Comment{Id: int64(comment.ID), Content: comment.Content, CreateDate: comment.CreatedAt.Format(time.Layout)}
		user, err := model.QueryUserById(int(comment.UserID))
		if err != nil {
			return res, err
		}
		now.User = common.User{Id: int64(user.ID), Name: user.Username, FollowCount: user.FollowCount, FollowerCount: user.FollowerCount}
		now.User.IsFollow = false // to modify later
		res = append(res, now)
	}
	return res, nil
}
