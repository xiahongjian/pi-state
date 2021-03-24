package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xiahongjian/pi-status/pkg/setting"
	"github.com/xiahongjian/pi-status/routers/api"
	"github.com/xiahongjian/pi-status/routers/tpl"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	api.InitApiRouter(engine)
	tpl.InitPageRouter(engine)

	return engine
}
