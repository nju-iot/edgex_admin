package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nju-iot/edgex_admin/handler"
	"github.com/nju-iot/edgex_admin/wrapper"
)

func registerRouter(r *gin.Engine) {
	r.GET("/ping", handler.Ping)
	// your code
	edgexRouter := r.Group("/edgex_admin")
	{
		edgexRouter.GET("/get_all_edgex", wrapper.JsonOutPutWrapper(handler.GetAllEdgex))
	}
}
