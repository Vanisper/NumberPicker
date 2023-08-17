package config

import (
	"context"
	"fmt"
	"number-picker/utils"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ConfigType struct {
	Category map[string][]int `toml:"category"`
}

var AppName = "NumberPicker"
var AppDescription = "彩票选号器"
var AppVersion = "0.0.1"
var AppConfigPath = utils.GetCurrPath() + "/configs"

func InitConfig(ctx context.Context) {
	// 初始化配置目录
	utils.MkDir(AppConfigPath)
	// 设置配置文件名和路径
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(AppConfigPath)
	// 监听
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件发生变化:", e)
		fmt.Println("配置文件发生变化:", viper.AllSettings())
		// 映射到结构体
		var config ConfigType
		if err := viper.Unmarshal(&config); err != nil {
			panic(err)
		}
		runtime.EventsEmit(ctx, "update:config", config)
	})
	viper.WatchConfig()

	// 设置默认值
	/*
		[分类配置]
		红球=01, 02, 03, 04, 05, 06, 07, 08, 09, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33
		篮球=01, 02, 03, 04, 05, 06, 07, 08, 09, 10, 11, 12, 13, 14, 15, 16
		金=01, 02, 09, 10, 23, 24
		木=11, 12, 13, 14, 22, 23
		水=05, 06, 19, 20, 27, 28
		火=07, 08, 15, 16, 29, 30
		土=03, 04, 17, 18, 25, 26
		单=01, 03, 05, 07, 09, 11, 13, 15, 17, 19, 23, 25, 27, 29, 31, 33
		双=02, 04, 06, 08, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32
		和单=03, 05, 07, 09, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32
		和双=04, 06, 08, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33
	*/
	viper.SetDefault("category.红球", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33})
	viper.SetDefault("category.篮球", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	viper.SetDefault("category.金", []int{1, 2, 9, 10, 23, 24})
	viper.SetDefault("category.木", []int{11, 12, 13, 14, 22, 23})
	viper.SetDefault("category.水", []int{5, 6, 19, 20, 27, 28})
	viper.SetDefault("category.火", []int{7, 8, 15, 16, 29, 30})
	viper.SetDefault("category.土", []int{3, 4, 17, 18, 25, 26})
	viper.SetDefault("category.单", []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 23, 25, 27, 29, 31, 33})
	viper.SetDefault("category.双", []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32})
	viper.SetDefault("category.和单", []int{3, 5, 7, 9, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32})
	viper.SetDefault("category.和双", []int{4, 6, 8, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33})

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		// 如果配置文件不存在，则创建一个新的配置文件
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = viper.SafeWriteConfig()
			if err != nil {
				panic(err)
			}
		} else {
			// 配置文件被找到，但产生了另外的错误
			panic(err)
		}
	}
	// 映射到结构体
	var config ConfigType
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	fmt.Println(config)
}

func GetConfig() ConfigType {
	// 映射到结构体
	var config ConfigType
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	return config
}
