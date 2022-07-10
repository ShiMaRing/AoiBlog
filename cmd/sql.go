package cmd

import (
	"Aoi/internal/sql2struct"
	"github.com/spf13/cobra"
	"log"
)

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql转结构体",
	Long:  "sql转结构体",
	Run: func(cmd *cobra.Command, args []string) {
		var dbInfo = sql2struct.DBInfo{
			DBType:   "mysql",
			Host:     "101.43.161.75:3306",
			UserName: "Aoi",
			Password: "123456",
			Charset:  "utf8",
		}
		model := sql2struct.NewDBModel(&dbInfo)
		err := model.Connect()
		if err != nil {
			log.Fatalln(err)
		}
		columns, err := model.GetColumns("aoi", "blog_tag")

		if err != nil {
			log.Fatalln(err)
		}
		template := sql2struct.NewStructTemplate()
		assemblyColumns := template.AssemblyColumns(columns)
		err = template.Generate("blog_tag", assemblyColumns)
		if err != nil {
			log.Fatalln(err)
		}
	},
}
