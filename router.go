package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nju-iot/edgex_admin/handlers"
	"github.com/nju-iot/edgex_admin/handlers/edgex"
	"github.com/nju-iot/edgex_admin/wrapper"
)

func registerRouter(r *gin.Engine) {
	r.GET("/ping", handlers.Ping)
	// your code

	edgexRouter := r.Group("/edgex_admin/edgex")
	{
		edgexRouter.GET("/search", wrapper.JSONOutPutWrapper(edgex.SearchEdgex))
		edgexRouter.POST("/create", wrapper.JSONOutPutWrapper(edgex.CreateEdgex))
		edgexRouter.POST("/delete", wrapper.JSONOutPutWrapper(edgex.DeleteEdgex))
		edgexRouter.POST("/follow", wrapper.JSONOutPutWrapper(edgex.FollowEdgex))
		edgexRouter.POST("/unfollow", wrapper.JSONOutPutWrapper(edgex.UnFollowEdgex))
	}
}
