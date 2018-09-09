package models

import (
	"theAmazingNotificator/app/common"
	"time"
)

type Post struct {
	Id              uint      `gorm:"primary_key"`
	AuthorID        uint      `gorm:"not null" json:"-"`
	Author          User      `gorm:"ForeignKey:AuthorID"`
	Title           string    `gorm:"not null"`
	Description     string    `gorm:"not null;type:text"`
	Comments        []Comment `gorm:"ForeignKey:PostID"`
	Votes           int       `gorm:"default:0"`
	Created         time.Time `gorm:"default:current_timestamp"`
	CommentQuantity int       `gorm:"default:0"`
}

func GetPostById(id uint) (Post, bool, error) {

	post := Post{}

	r := common.GetDatabase()

	r = r.Where("id = ?", id).Preload("Author").First(&post)
	if r.RecordNotFound() {
		return post, false, nil
	}

	if r.Error != nil {
		return post, true, r.Error
	}

	return post, true, nil
}
