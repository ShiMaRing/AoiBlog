package dao

import (
	"Aoi/internal/model"
	sql2struct2 "Aoi/internal/tools/sql2struct"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// CreateTables 对表进行创建
func CreateTables() error {
	dbConn := GetDB()
	db := dbConn.DBConnection
	err := db.AutoMigrate(&model.Article{})
	if err != nil {
		return err
	}
	return nil
}

// GetDB 返回一个DBModel对象
func GetDB() *sql2struct2.DBModel {
	var err error
	var dbInfo = sql2struct2.DBInfo{
		DBType:   "mysql",
		Host:     "101.43.161.75:3306",
		UserName: "Aoi",
		Password: "123456",
		Charset:  "utf8",
	}
	dbModel := sql2struct2.NewDBModel(&dbInfo)
	dsn := "%s:%s@tcp(%s)/aoi?charset=%s&parseTime=True&loc=Local"
	info := dbModel.DBInfo
	dsn = fmt.Sprintf(dsn,
		info.UserName,
		info.Password,
		info.Host,
		info.Charset)
	dbModel.DBConnection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return dbModel
}
