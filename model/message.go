package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FromID  int64
	ToID    int64
	Content string
}

func (Message) TableName() string {
	return "messages"
}

func AddMessage(message *Message) error {
	return db.Create(message).Error
}

func QueryMessageById(fromID, toID, time int64) []*Message {
	var messages []*Message
	db.Where("(from_id = ? and to_id = ? or from_id = ? and to_id = ?) and created_at > ?", fromID, toID, toID, fromID, time).Order("created_at").Find(&messages)
	fmt.Println(time)
	return messages
}
