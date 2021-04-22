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
		edgexRouter.GET("/search", wrapper.JsonOutPutWrapper(edgex.SearchEdgex))
		edgexRouter.POST("/create", wrapper.JsonOutPutWrapper(edgex.CreateEdgex))
		edgexRouter.POST("/delete", wrapper.JsonOutPutWrapper(edgex.DeleteEdgex))
		edgexRouter.POST("/follow", wrapper.JsonOutPutWrapper(edgex.FollowEdgex))
		edgexRouter.POST("/unfollow", wrapper.JsonOutPutWrapper(edgex.UnFollowEdgex))
	}
}
