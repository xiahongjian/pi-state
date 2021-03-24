package api

import "github.com/gin-gonic/gin"

func InitApiRouter(engine *gin.Engine) {
	group := engine.Group("/api")
	{
		group.GET("/state", SummaryInfo)
	}
}
