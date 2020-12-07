package main

import (
	"ginServer/config"
	"ginServer/handle"
	"ginServer/logger"
	"github.com/gin-gonic/gin"
)

var Port = ":" + config.Cfg.ServerPort

func setupRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	logger.Log.Infof("get rpcport1:%s, port2:%s",
		config.Cfg.VideoReportPort,
		config.Cfg.VideoRecommPort)

	// 注册rpc videoReport 服务
	videoReportClientServer := handle.VideoReportClientServer{
		Engine: router,
		Client: handle.InitClient(":" + config.Cfg.VideoReportPort),
	}
	// 注册rpc recomm推荐服务
	recommClientServer := handle.RecommClientServer{
		Engine: router,
		Client: handle.InitRecommClient(":" + config.Cfg.VideoRecommPort),
	}

	//http://localhost:8006/ping?key=hahah
	router.GET("/ping", handle.Ping)

	//curl -X POST http://localhost:8006/postdata -d 'first=hawrk&second=chen'
	router.POST("/postdata", handle.PostData)

	// 获取 统计报表数据
	router.POST("/getReportData", videoReportClientServer.GetVideoReportData)
	// 获取推荐视频数据
	router.POST("/getRecommVideoLists", recommClientServer.GetRecommVideoList)

	return router
}

func main() {
	logger.Log.Info(">>>>>>>>>>>> service start >>>>>>>>>>>>>>>")
	logger.Log.Debugf("get log level: %v", config.Cfg.LogLevel)

	r := setupRouter()
	r.Run(Port)

	// 强刷缓冲区
	defer logger.Log.Sync()
}
