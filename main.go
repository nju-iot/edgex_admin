package main

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/nju-iot/edgex_admin/caller"
	"github.com/nju-iot/edgex_admin/config"
	"github.com/nju-iot/edgex_admin/cronloader"
	"github.com/nju-iot/edgex_admin/logs"
	"go.uber.org/zap"
)

func main() {

	config.LoadConfig()
	logs.InitLogs()
	caller.InitClient()
	cronloader.InitCronLoader()

	gin.SetMode(config.Server.RunMode)

	r := gin.New()
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))

	registerRouter(r)

	_ = r.Run(config.Server.Port)
}
