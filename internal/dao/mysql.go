package dao

import (
	"Aoi/global"
	"Aoi/internal/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// CreateTables 对表进行创建
func CreateTables() error {
	db, _ := GetDB()
	err := db.AutoMigrate(&model.Auth{})
	if err != nil {
		return err
	}
	return nil
}

// GetDB 返回一个DBModel对象
func GetDB() (*gorm.DB, error) {
	var err error
	var databaseSetting = global.DatabaseSetting

	//根据配置文件解析
	dsn := "%s:%s@tcp(%s)/aoi?charset=%s&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.Charset)
	//配置日志参数
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,          // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, err
	}

	s, _ := db.DB()
	s.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	s.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil
}
