package tasks

import (
	"encoding/json"
	"theAmazingNotificator/app/models"
)

type CommentCommentNotificationTask struct {
	Queue           string
	FatherCommentID uint
	CommentID       uint
	Type            uint
}

func NewCommentCommentNotificationTask(fatherCommentID uint, commentID uint) CommentCommentNotificationTask {

	return CommentCommentNotificationTask{
		Queue:           "comment_comment_notification_queue",
		FatherCommentID: fatherCommentID,
		CommentID:       commentID,
		Type:            models.CommentComment,
	}
}

func (t CommentCommentNotificationTask) GetMessageBytes() ([]byte, error) {

	data, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (t CommentCommentNotificationTask) GetQueue() (queueName string) {
	return t.Queue
}
