package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xiahongjian/pi-status/pkg/setting"
	"github.com/xiahongjian/pi-status/routers/api"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	group := router.Group("/api")
	{
		group.GET("/state", api.SummaryInfo)
	}

	return router
}
