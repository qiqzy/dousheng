package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Name     string `gorm:"unique;DEFAULT:'未定义'"`
	// 关注
	FollowCount int64 `gorm:"DEFAULT:0"`
	// 粉丝
	FollowerCount int64 `gorm:"DEFAULT:0"`
}

func (User) TableName() string {
	return "users"
}

func AddUser(user User) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func QueryUserById(id int) (*User, error) {
	user := &User{}
	result := db.Where("id = ?", id).Find(user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		user = nil
	}
	return user, nil
}

func QueryUserByName(name string) (*User, error) {
	user := &User{}
	result := db.Where("username = ?", name).Find(user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		user = nil
	}
	return user, nil
}

func UpdateUser(user User) error {
	if err := db.Save(user).Error; err != nil {
		return err
	}
	return nil
}
func DeleteUser(name string) error {
	if err := db.Where("username = ?", name).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}
