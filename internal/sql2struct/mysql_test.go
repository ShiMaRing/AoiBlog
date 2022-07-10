package sql2struct

import (
	"fmt"
	"testing"
)

//测试通过
func TestDBModel_GetColumns(t *testing.T) {
	var dbInfo = DBInfo{
		DBType:   "mysql",
		Host:     "101.43.161.75:3306",
		UserName: "Aoi",
		Password: "123456",
		Charset:  "utf8",
	}
	model := NewDBModel(&dbInfo)
	err := model.Connect()
	if err != nil {
		t.Error(err)
	}
	columns, err := model.GetColumns("aoi", "blog_tag")
	if err != nil {
		t.Error(err)
	}
	for _, column := range columns {
		fmt.Println(column)
	}
}
