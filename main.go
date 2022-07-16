package main

import (
	_ "Aoi/docs"
	"Aoi/global"
	"Aoi/internal/dao"
	"Aoi/internal/routers"
	"Aoi/pkg/logger"
	"Aoi/pkg/setting"
	"Aoi/pkg/tracer"
	"context"
	"flag"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	signal2 "os/signal"
	"strings"
	"syscall"
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
	err = setupTracer()
	if err != nil {
		log.Fatalln(err)
	}
	err = setupFlag() //添加命令行参数绑定
	if err != nil {
		log.Fatalln(err)
	}

}

//@title AoiBlog
//@version 1.0
//@description Aoi博客
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
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Print(err)
		}
	}()

	var signal = make(chan os.Signal)
	signal2.Notify(signal, syscall.SIGINT, syscall.SIGTERM) //捕获并往signal中放置数据
	<-signal

	log.Println("serve is shouting down")
	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	if err := server.Shutdown(timeout); err != nil {
		log.Println("close with error ", err)
	}
	log.Println("see you !")
}

func setupSetting() error {
	var err error
	set = setting.NewSetting(strings.Split("config", ",")...)
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
	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	err = set.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second //7200秒

	err = set.ReadSection("Email", &global.EmailSetting)
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

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("aoi", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}

var (
	port    string
	runMode string
	config  string
)

func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "config/", "配置文件路径")
	flag.Parse()
	return nil
}
