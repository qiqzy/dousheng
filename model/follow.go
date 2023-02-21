package model

import "gorm.io/gorm"

// Follow 关注
type Follow struct {
	gorm.Model
	// 粉丝 Follower 关注 Follow
	FollowerID int64
	// 被关注者
	FollowID int64
}

func (Follow) TableName() string {
	return "follows"
}

func AddFollow(followerId int64, followId int64) error {
	result := db.Create(&Follow{FollowerID: followerId, FollowID: followId})
	return result.Error
}

func DeleteFollow(followerId int64, followId int64) error {
	return db.Where("follower_id = ? and follow_id = ?", followerId, followId).Delete(&Follow{}).Error
}

func QueryFollowByFollowerId(followerId int64) ([]*Follow, error) {
	var res []*Follow
	result := db.Where("follower_id = ?", followerId).Find(&res)
	return res, result.Error
}

func QueryFollowByFollowId(followId int64) ([]*Follow, error) {
	var res []*Follow
	result := db.Where("follow_id = ?", followId).Find(&res)
	return res, result.Error
}

func QueryFollowByFollowerIdFollowId(followerId int64, followId int64) (bool, error) {
	result := db.Where("follower_id = ? and follow_id = ?", followerId, followId).Find(&Follow{})
	res := false
	if result.RowsAffected > 0 {
		res = true
	}
	return res, result.Error
}
