package models

import (
	"github.com/jinzhu/gorm"
	"theAmazingNotificator/app/common"
)

type User struct {
	gorm.Model
	GUID                 string `gorm:"type:char(20);unique_index:idx_unique_guid_object" json:"ID"`
	Name                 string `gorm:"null"`
	LastName             string `gorm:"null"`
	Password             string `gorm:"null" json:"-"`
	Email                string
	GoogleID             string `gorm:"null" json:"-"`
	Phone                string `gorm:"null"`
	PasswordRecoveryCode string `gorm:"null" json:"-"`
	RoleID               uint   `gorm:"not null" json:"-"`
	Enabled              bool   `gorm:"default:true"`
}

func GetUserById(id uint) (user User, found bool, err error) {

	user = User{}

	r := common.GetDatabase()

	r = r.Unscoped().Where("id = ?", id).First(&user)
	if r.RecordNotFound() {
		return user, false, nil
	}

	if r.Error != nil {
		return user, true, r.Error
	}

	return user, true, nil
}
