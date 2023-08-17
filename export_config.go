package main

import (
	config "number-picker/configs"
)

func (a *App) ExportConfig() config.ConfigType {
	// 返回viper配置
	return config.GetConfig()
}
