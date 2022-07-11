package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Common               //嵌入common字段
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"coverImageUrl"`
	State         uint8  `json:"state"`
	Tags          []*Tag `json:"-" gorm:"many2many:article_tag"`
}

func (a Article) TableName() string {
	return "article"
}
