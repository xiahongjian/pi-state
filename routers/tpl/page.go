package tpl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiahongjian/pi-status/models"
	"github.com/xiahongjian/pi-status/service"
)

func InitPageRouter(engine *gin.Engine) {
	engine.Group("/")
	{
		engine.GET("/index.html", Index)
		engine.GET("/", func(c *gin.Context) {
			c.Request.URL.Path = "/index.html"
			engine.HandleContext(c)
		})
	}
}

func Index(c *gin.Context) {
	service := &service.StateService{}
	c.HTML(http.StatusOK, "index.tpl", models.R{
		Data: service.GetState(),
	})
}
