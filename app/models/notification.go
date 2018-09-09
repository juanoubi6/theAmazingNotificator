package models

import (
	"theAmazingNotificator/app/common"
	"time"
)

type Notification struct {
	ID                 uint             `gorm:"primary_key"`
	Message            string           `gorm:"not null"`
	UserID             uint             `gorm:"not null"`
	Created            time.Time        `gorm:"default:current_timestamp"`
	NotificationType   NotificationType `gorm:"ForeignKey:NotificationTypeID"`
	NotificationTypeID uint
}

func (notificationData *Notification) Save() error {

	err := common.GetDatabase().Create(notificationData).Error
	if err != nil {
		return err
	}

	return nil

}

func (notificationData *Notification) Delete() error {

	err := common.GetDatabase().Delete(notificationData).Error
	if err != nil {
		return err
	}

	return nil

}
