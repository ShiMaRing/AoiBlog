package sql2struct

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// TableColumn 提供数据库字段的查找
type TableColumn struct {
	ColumnName    string `gorm:"column:COLUMN_NAME"`    //列名
	DataType      string `gorm:"column:DATA_TYPE"`      //数据类型
	IsNullable    string `gorm:"column:IS_NULLABLE"`    //是否允许null值
	ColumnKey     string `gorm:"column:COLUMN_KEY"`     //是什么类型的键
	ColumnType    string `gorm:"column:COLUMN_TYPE"`    //列数据类型
	ColumnComment string `gorm:"column:COLUMN_COMMENT"` //注释
}

// DBInfo 填写相关连接信息
type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

// DBModel 提供数据库连接管理
type DBModel struct {
	DBConnection *gorm.DB
	DBInfo       *DBInfo
}

// NewDBModel 创建新的数据库连接
func NewDBModel(DBInfo *DBInfo) *DBModel {
	return &DBModel{DBInfo: DBInfo}
}

// Connect 连接数据库，完成此方法后db对象中应当持有连接对象
func (model *DBModel) Connect() error {
	var err error
	dsn := "%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local"
	info := model.DBInfo
	dsn = fmt.Sprintf(dsn,
		info.UserName,
		info.Password,
		info.Host,
		info.Charset)
	model.DBConnection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

//gorm能够对指针切片进行注入
func (model *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	var tableColumns []*TableColumn
	connection := model.DBConnection
	err := connection.Table("COLUMNS").Where("table_schema=?", dbName).Where("table_name=?", tableName).Find(&tableColumns).Error
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return tableColumns, nil
}

var DBTypeToStructType = map[string]string{
	"int":       "int32",
	"tinyint":   "int8",
	"smallint":  "int",
	"mediumint": "int64",
	"bigint":    "int64",
	"bit":       "int",
	"bool":      "bool",
	"enum":      "string",
	"set":       "string",
	"varchar":   "string",
}
