package utils

// viper
// 用于读取config  yml
import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// InitConfig
// 用于读取config里面的配置
func InitConfig() {
	viper.AddConfigPath("config")
	viper.SetConfigName("app")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalln("Init Config :", err)
		return
	}

	fmt.Println("config app inited ...")
}