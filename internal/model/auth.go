package model

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	AppKey    string `json:"appKey,omitempty" gorm:"column:app_key"`
	AppSecret string `json:"appSecret,omitempty" gorm:"column:app_secret"`
}

func (a Auth) TableName() string {
	//TODO implement me
	return "auth_blog"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	err := db.Where("app_key=? and app_secret=?", a.AppKey, a.AppSecret).First(&auth).Error
	if err != nil {
		return auth, err
	}
	return auth, nil
}

func (a Auth) Add(db *gorm.DB) error {
	return db.Create(&a).Error
}
