package model

import (
	"Aoi/global"
	"fmt"
	"gorm.io/gorm"
)

// Tag tag与article之间应当为多对多关系
type Tag struct {
	gorm.Model
	Common
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) String() string {
	//TODO implement me
	return fmt.Sprintf("%d %s %d", t.ID, t.Name, t.State)
}

func (t Tag) TableName() string {
	return "tag"
}

// Count tag 和文章是多对多的关系，该方法为由中间关系表查询相关名字的tag数量
// Count 根据传入的tag名称获取tag数量
// tag的名称与id都是唯一的，根据名称查询tag的使用数量
func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int64
	db.Table("article_tag").Where("tag_id=(?)", db.Model(t).
		Select("id").Where("name=?", t.Name).Where("state=?", t.State)).Count(&count)
	return int(count), nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Scopes(Paginate(pageOffset, pageSize))
	} else {
		pageOffset = 0
		pageSize = global.AppSetting.DefaultPageSize
	}
	stmt := db.Model(&Tag{}).Where("state=?", t.State).Find(&tags).Statement

	err = stmt.Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func Paginate(pageOffset, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pageOffset).Limit(pageSize)
	}
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB) error {
	return db.Model(&t).Select("name", "state", "updated_by").Updates(t).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Delete(&t).Error
}
