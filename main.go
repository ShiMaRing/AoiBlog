package main

import (
	"Aoi/global"
	"Aoi/internal/routers"
	"Aoi/pkg/setting"
	"log"
	"net/http"
	"time"
)

var name string
var set *setting.Setting

func main() {
	err := setupSetting()
	if err != nil {
		log.Fatalln(err)
	}
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
