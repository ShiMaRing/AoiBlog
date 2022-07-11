package cmd

import (
	sql2struct2 "Aoi/internal/tools/sql2struct"
	"github.com/spf13/cobra"
	"log"
)

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql转结构体",
	Long:  "sql转结构体",
	Run: func(cmd *cobra.Command, args []string) {
		var dbInfo = sql2struct2.DBInfo{
			DBType:   "mysql",
			Host:     "101.43.161.75:3306",
			UserName: "Aoi",
			Password: "123456",
			Charset:  "utf8",
		}
		model := sql2struct2.NewDBModel(&dbInfo)
		err := model.Connect()
		if err != nil {
			log.Fatalln(err)
		}
		columns, err := model.GetColumns("aoi", "blog_tag")

		if err != nil {
			log.Fatalln(err)
		}
		template := sql2struct2.NewStructTemplate()
		assemblyColumns := template.AssemblyColumns(columns)
		err = template.Generate("blog_tag", assemblyColumns)
		if err != nil {
			log.Fatalln(err)
		}
	},
}
