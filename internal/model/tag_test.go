package model

import (
	"Aoi/pkg/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestTagCount(t *testing.T) {
	var tag = Tag{
		Name:  "aa",
		State: 1,
	}
	var db, _ = dbs()
	count, _ := tag.Count(db)
	fmt.Println(count)

}

func TestTagList(t *testing.T) {
	var tag = Tag{
		Name:  "aa",
		State: 1,
	}
	var db, _ = dbs()
	count, _ := tag.List(db, 0, 1)
	fmt.Println(count)
}

func dbs() (*gorm.DB, error) {
	var err error
	var dataset = setting.DatabaseSettingS{
		DBType:   "mysql",
		UserName: "Aoi",
		Password: "123456",
		Host:     "101.43.161.75:3306",
		DBName:   "aoi",
		Charset:  "utf8",
	}
	var databaseSetting = &dataset
	//根据配置文件解析
	dsn := "%s:%s@tcp(%s)/aoi?charset=%s&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.Charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestTagCreate(t *testing.T) {
	var tag = Tag{
		Name:  "mm",
		State: 1,
	}
	var db, _ = dbs()
	err := tag.Create(db)
	t.Log(err)

}

func TestTagUpdate(t *testing.T) {
	var tag = Tag{
		Model: gorm.Model{ID: 4},
		Name:  "qqqqq",
		State: 0,
	}
	var db, _ = dbs()
	err := tag.Update(db)
	t.Log(err)

}
