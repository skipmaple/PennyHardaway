package models

import (
	"PennyHardway/pkg/logging"
	"github.com/jinzhu/gorm"
)

type Auth struct {
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) (bool, error) {
	var auth Auth
	//logging.Debug(db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error)
	logging.Info("****", db == nil)
	err := db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Error(err)
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}
