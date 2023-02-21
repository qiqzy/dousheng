package main

import (
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	//go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	//加载配置文件
	config.InitConfig()

	//创建数据库连接
	model.InitDb()

	gin.SetMode(viper.GetString("server.run_mode"))
	r.Run(viper.GetString("server.addr"))
	//r.Run()
}
