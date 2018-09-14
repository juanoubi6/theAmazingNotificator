package models

import (
	"theAmazingNotificator/app/common"
	"time"
)

type Notification struct {
	ID                 uint             `gorm:"primary_key"`
	Message            string           `gorm:"not null"`
	UserID             uint             `gorm:"not null" json:"-"`
	Created            time.Time        `gorm:"default:current_timestamp"`
	NotificationType   NotificationType `gorm:"ForeignKey:NotificationTypeID"`
	NotificationTypeID uint             `json:"-"`
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

func GetAllNotifications(userID uint, limit int, offset int) ([]Notification, int, error) {

	var notifications []Notification
	var quantity int

	//Get notifications
	r := common.GetDatabase().Preload("NotificationType").Where("user_id = ?", userID).Limit(limit).Offset(offset).Order("created desc").Find(&notifications)
	if r.Error != nil {
		return notifications, 0, r.Error
	}

	//Get posts quantity
	r = common.GetDatabase().Table("notifications").Where("user_id = ?", userID).Count(&quantity)
	if r.Error != nil {
		return notifications, 0, r.Error
	}

	return notifications, quantity, nil

}
