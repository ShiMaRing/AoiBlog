package model

import "gorm.io/gorm"

// Tag tag与article之间应当为多对多关系
type Tag struct {
	gorm.Model
	Common
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "tag"
}
