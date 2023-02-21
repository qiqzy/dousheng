package model

import "gorm.io/gorm"

// Favorite 点赞
type Favorite struct {
	gorm.Model
	UserID  int64
	VideoID int64
}

func (Favorite) TableName() string {
	return "favorites"
}

func AddFavorite(userId int64, videoId int64) error {
	if err := db.Create(&Favorite{UserID: userId, VideoID: videoId}).Error; err != nil {
		return err
	}
	return nil
}

func QueryFavoriteByuserId(userId int64) ([]*Favorite, error) {
	var res []*Favorite
	result := db.Where("user_id = ?", userId).Find(&res)
	return res, result.Error
}

func QueryFavoriteByUserIdVideoId(userId int64, videoId int64) (bool, error) {
	result := db.Where("user_id = ? and video_id = ?", userId, videoId).Find(&Favorite{})
	res := false
	if result.RowsAffected > 0 {
		res = true
	}
	return res, result.Error
}

func DeleteFavorite(userId int64, videoId int64) error {
	if err := db.Where("user_id = ? and video_id = ?", userId, videoId).Delete(&Favorite{}).Error; err != nil {
		return err
	}
	return nil
}
