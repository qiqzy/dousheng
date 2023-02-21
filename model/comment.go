package model

import "gorm.io/gorm"

// Comment 评论
type Comment struct {
	gorm.Model
	UserID  int64
	VideoID int64
	Content string
}

func (Comment) TableName() string {
	return "comments"
}

func AddComment(comment *Comment) error {
	return db.Create(comment).Error
}

func QueryCommentById(id int64) (*Comment, error) {
	comment := &Comment{}
	result := db.Where("id = ?", id).Find(comment)
	if result.RowsAffected == 0 {
		comment = nil
	}
	return comment, result.Error
}

func DeleteCommentById(id int64) error {
	return db.Where("id = ?", id).Delete(&Comment{}).Error
}

func QueryCommentByVideoId(videoId int64) ([]*Comment, error) {
	var ans []*Comment
	result := db.Where("video_id = ?", videoId).Order("updated_at desc").Find(&ans)
	err := result.Error
	return ans, err
}
