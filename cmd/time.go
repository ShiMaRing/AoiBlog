package cmd

import (
	"Aoi/internal/tools/timer"
	"github.com/spf13/cobra"
	"log"
	"strings"
	"time"
)

//提供time命令

var timeCmd = &cobra.Command{
	Use:     "time",
	Aliases: nil,
	Short:   "时间格式处理",
	Long:    "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取系统当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		now := timer.GetNowTime()
		log.Printf("当前时间为：%s, %d", now.Format("2006-01-02 15:04:05"), now.Unix())
	},
}
var calTimeCmd = &cobra.Command{
	Use:   "cal",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var cur time.Time
		var layout = "2006-01-02 15:04:05"
		var pre = layout
		if calTime == "" {
			//说明用户没有输入，此时默认采用当前时间
			cur = timer.GetNowTime()
		} else {
			var err error
			//根据用户的输入方式判断时间
			count := strings.Count(calTime, " ")
			if count == 0 {
				layout = "2006-01-02"
			} else if count == 1 {
				layout = "2006-01-02 15:04:05"
			}
			cur, err = time.Parse(layout, calTime)
			if err != nil {
				log.Fatalf("时间格式错误 %v  %s", err, calTime)
				return
			}
		}
		getCalTime, err := timer.GetCalTime(cur, duration)
		if err != nil {
			log.Fatalf("时间计算错误 err:%v", err)
		}
		log.Printf("计算结果：%s\n", getCalTime.Format(pre))
	},
}
var calTime string
var duration string

func init() {
	timeCmd.AddCommand(nowCmd)
	timeCmd.AddCommand(calTimeCmd)

	calTimeCmd.Flags().StringVarP(&calTime, "calculate", "c", "", "输入需要计算的时间，不输入默认采用当前时间")
	calTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "输入时间间隔")
}
