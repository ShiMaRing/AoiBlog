package cmd

import (
	"Aoi/internal/word"
	"github.com/spf13/cobra"
	"log"
)

const (
	ModeUpper                      = iota + 1 //全转大写
	ModeLower                                 //全转小写
	ModeUnderscoreToUpperCamelCase            //下划线转驼峰大写
	ModeUnderscoreToLowerCamelcase            //下划线转驼峰小写
	ModeCamelCaseToUnderscore                 //驼峰转下划线
)

var str string
var mode int8

var desc = `
1:全部单词转为大写,
2:全部单词转为小写,
3:下画线单词转为大写驼峰单词,
4:下画线单词转为小写驼峰单词,
5:驼峰单词转为下画线单词,
`

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转化",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		//在这里面定义具体的执行方法
		var content string
		switch mode {
		case ModeLower:
			content = word.ToLower(str)
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)

		case ModeUnderscoreToLowerCamelcase:
			content = word.UnderscoreToLowerCameCase(str)

		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseUnderscore(str)
		default:
			log.Fatalf("暂不支持该模式转化")
		}
		log.Printf("result : %v", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "选择模式转化方式")
}
