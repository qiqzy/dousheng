package service

import (
	"errors"
	"fmt"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/model"
	"os/exec"
)

func Feed(latest_time int64) ([]common.Video, error) { //to add latest_time mode
	var res []common.Video
	videos, err := model.QueryVideoByTime()
	if err != nil {
		return res, err
	}
	n := len(videos)
	for i := 0; i < n; i++ {
		now := videos[i]
		aVideo := common.Video{Id: int64(now.ID), PlayUrl: now.PlayUrl, CoverUrl: now.CoverUrl, FavoriteCount: now.FavoriteCount, CommentCount: now.CommentCount}
		user, err := model.QueryUserById(int(now.AuthorID))
		if err != nil {
			return nil, err
		}
		if user == nil {
			return nil, errors.New("The video can't find the author")
		}
		aVideo.Author = common.User{Id: int64(user.ID), Name: user.Name, FollowCount: user.FollowCount, FollowerCount: user.FollowerCount}
		aVideo.Author.IsFollow = false
		aVideo.IsFavorite = false
		res = append(res, aVideo)
	}
	return res, nil
}

func Publish(authorID int64, title string, play_url string, cover_url string) error {
	video := model.Video{AuthorID: authorID, Title: title, PlayUrl: play_url, CoverUrl: cover_url}
	err := model.AddVideo(video)
	return err
}

func List(authorID int64) ([]common.Video, error) {
	var res []common.Video
	videos, err := model.QueryVideoByAuthor(authorID)
	if err != nil {
		return res, err
	}
	n := len(videos)
	for i := 0; i < n; i++ {
		now := videos[i]
		aVideo := common.Video{Id: int64(now.ID), PlayUrl: now.PlayUrl, CoverUrl: now.CoverUrl, FavoriteCount: now.FavoriteCount, CommentCount: now.CommentCount}
		user, err := model.QueryUserById(int(now.AuthorID))
		if err != nil {
			return nil, err
		}
		if user == nil {
			return nil, errors.New("The video can't find the author")
		}
		aVideo.Author = common.User{Id: int64(user.ID), Name: user.Name, FollowCount: user.FollowCount, FollowerCount: user.FollowerCount}
		aVideo.Author.IsFollow = true
		aVideo.IsFavorite = true
		res = append(res, aVideo)
	}
	return res, nil
}

func Generate(videoPath string, videoName string, coverName string) error {
	cmd := exec.Command("ffmpeg", "-i", videoPath+videoName, "-ss", "1", "-f",
		"image2", "-frames:v", "1", videoPath+coverName)
	fmt.Println(cmd.String())

	err := cmd.Run()
	return err
}
