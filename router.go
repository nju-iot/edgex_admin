package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nju-iot/edgex_admin/handler"
)

func registerRouter(r *gin.Engine) {
	r.GET("/ping", handler.Ping)
	// your code
	edgexRouter := r.GET("/edgex_admin")
	{
		edgexRouter.GET("/")

	}
}
