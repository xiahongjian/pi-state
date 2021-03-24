package main

import (
	"fmt"
	"net/http"

	"github.com/xiahongjian/pi-status/pkg/setting"
	"github.com/xiahongjian/pi-status/routers"
)

func main() {
	router := routers.InitRouter()
	router.LoadHTMLGlob(setting.TemplatePath)
	router.StaticFS("/static", http.Dir(setting.StaticPath))

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
