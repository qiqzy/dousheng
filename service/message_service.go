package service

import (
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/model"
)

func MessageAdd(fromId int64, toId int64, content string) error {
	message := &model.Message{FromID: fromId, ToID: toId, Content: content}
	err := model.AddMessage(message)
	return err
}

func MessageList(fromID int64, toID int64, time int64) []common.Message {
	messages := model.QueryMessageById(fromID, toID, time)
	res := make([]common.Message, len(messages))
	for index, message := range messages {
		res[index] = common.Message{
			Id:         int64(message.ID),
			ToUserId:   message.ToID,
			FromUserId: message.FromID,
			Content:    message.Content,
		}
	}
	return res
}
