package models

import (
	"theAmazingNotificator/app/common"
	"time"
)

type Comment struct {
	Id              uint      `gorm:"primary_key"`
	Message         string    `gorm:"not null"`
	AuthorID        uint      `gorm:"not null" json:"-"`
	Author          User      `gorm:"ForeignKey:AuthorID"`
	Votes           int       `gorm:"default:0"`
	Father          uint      `gorm:"default:0" json:"-"`
	PostID          uint      `gorm:"not null" json:"-"`
	Comments        []Comment `gorm:"ForeignKey:Father"`
	Created         time.Time `gorm:"default:current_timestamp"`
	CommentQuantity int       `gorm:"default:0"`
}

func GetCommentById(id uint) (Comment, bool, error) {

	comment := Comment{}

	r := common.GetDatabase()

	r = r.Where("id = ?", id).Preload("Author").First(&comment)
	if r.RecordNotFound() {
		return comment, false, nil
	}

	if r.Error != nil {
		return comment, true, r.Error
	}

	return comment, true, nil
}

