package global

import (
	"Aoi/pkg/logger"
	"Aoi/pkg/setting"
)

// Settings 定义全局变量给各个程序使用
var AppSetting *setting.AppSettingS
var ServerSetting *setting.ServerSettingS
var DatabaseSetting *setting.DatabaseSettingS
var Logger *logger.Logger
var JWTSetting *setting.JWTSettingS
var EmailSetting *setting.EmailSettingS
