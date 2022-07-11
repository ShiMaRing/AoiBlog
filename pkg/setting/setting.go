package setting

import (
	"github.com/spf13/viper"
	"log"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() *Setting {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("configs/")
	v.AddConfigPath("./configs/")

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("read config fail with %v \n", err)
		return nil
	}
	return &Setting{vp: v}
}
