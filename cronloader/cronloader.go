package cronloader

import (
	"github.com/jasonlvhit/gocron"
	"github.com/nju-iot/edgex_admin/utils"
)

func InitCronLoader() {
	gocron.Every(10).Seconds().Do(updateEdgexStatus)
	go func() {
		defer func() {
			utils.RecoverPanic()
		}()
		<-gocron.Start()
	}()
}
