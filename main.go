package main

import (
	"Aoi/global"
	"Aoi/internal/dao"
	"Aoi/internal/routers"
	"Aoi/pkg/logger"
	"Aoi/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

var name string
var set *setting.Setting

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalln(err)
	}
	err = setupDB()
	if err != nil {
		log.Fatalln(err)
	}
	_ = setupLogger()

}

func main() {

	var servSetting = global.ServerSetting
	router := routers.NewRouter()
	server := &http.Server{
		Addr:              ":" + servSetting.HttpPort,
		Handler:           router,
		ReadTimeout:       servSetting.ReadTimeout,
		ReadHeaderTimeout: servSetting.ReadTimeout,
		WriteTimeout:      servSetting.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Print(err)
	}
}

func setupSetting() error {
	var err error
	set = setting.NewSetting()
	err = set.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	//乘以秒数
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	err = set.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupDB() error {
	var err error
	global.Db, err = dao.GetDB()
	if err != nil {
		return err
	}
	return nil
}
func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   100,
		MaxAge:    10,
		LocalTime: true,
		Compress:  false,
	}, "", log.LstdFlags).WithCaller(0) //表示前缀出现在logger header尾部
	return nil
}
