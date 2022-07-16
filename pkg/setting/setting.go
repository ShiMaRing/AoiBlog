package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type Setting struct {
	vp *viper.Viper
}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig() //监控文件改变
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection() //重载新的配置文件
		})
	}()
}

func NewSetting(configs ...string) *Setting {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	for _, config := range configs {
		if config != "" {
			v.AddConfigPath(config)
		}
	}

	v.AddConfigPath("./configs/")

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("read config fail with %v \n", err)
		return nil
	}
	setting := &Setting{vp: v}
	setting.WatchSettingChange()
	return setting
}
