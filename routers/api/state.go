package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiahongjian/pi-status/models"
	"github.com/xiahongjian/pi-status/service"
)

func SummaryInfo(c *gin.Context) {
	service := &service.StateService{}
	c.JSON(http.StatusOK, models.R{
		Success: true,
		Data:    service.GetState(),
	})
}
