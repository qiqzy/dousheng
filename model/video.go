package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Title         string
	AuthorID      int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64 `gorm:"DEFAULT:0"`
	CommentCount  int64 `gorm:"DEFAULT:0"`
}

func (Video) TableName() string {
	return "videos"
}

func AddVideo(video Video) error {
	if err := db.Create(&video).Error; err != nil {
		return err
	}
	return nil
}

func QueryVideoByTime() ([]*Video, error) {
	var ans []*Video
	result := db.Limit(30).Order("updated_at desc").Find(&ans)
	err := result.Error
	return ans, err
}

func QueryVideoByAuthor(authorID int64) ([]*Video, error) {
	var ans []*Video
	result := db.Where("author_id = ?", authorID).Find(&ans)
	return ans, result.Error
}

func QueryVideoById(id int64) (*Video, error) {
	video := &Video{}
	result := db.Where("id = ?", id).Find(video)
	if result.RowsAffected == 0 {
		video = nil
	}
	return video, result.Error
}

func UpdateVideo(video Video) error {
	if err := db.Save(video).Error; err != nil {
		return err
	}
	return nil
}
