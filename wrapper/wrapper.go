package wrapper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nju-iot/edgex_admin/logs"
	"github.com/nju-iot/edgex_admin/resp"
)

func JsonOutPutWrapper(f func(*gin.Context) *JsonOutput) func(c *gin.Context) {
	return func(c *gin.Context) {
		var output *JsonOutput

		logs.Info("[wraper-request] url=%s, header=%v, body=%v",
			c.Request.URL, c.Request.Header, c.Request.Body)

		start := time.Now()

		defer func() {
			if tErr := recover(); tErr != nil {
				const size = 64 << 10
				buffer := make([]byte, size)
				buffer = buffer[:runtime.Stack(buffer, false)]
				logs.Error("[wrapper-panic] error=%v, stack=%s", tErr, buffer)

				rsp := resp.NewStdResponse(resp.RESP_CODE_SEVER_EXCEPTION, nil)
				output = NewJsonOutput(c, http.StatusInternalServerError, rsp)
			}
			if output == nil {
				logs.Error("[wraper-output-empty] output is empty!")
				rsp := resp.NewStdResponse(resp.RESP_CODE_SEVER_EXCEPTION, nil)
				output = NewJsonOutput(c, http.StatusInternalServerError, rsp)
			}

			output.Write()

			userTime := time.Since(start).Nanoseconds() / 1000
			logs.Info("[wraper-response] useTime=%d, status=%d, resp=%s",
				userTime, output.HttpStatus, GetMarshalStr(output.Resp))
		}()

		output = f(c)
	}
}

func GetMarshalStr(obj interface{}) string {
	vi := reflect.ValueOf(obj)
	if vi.Kind() == reflect.Ptr && vi.IsNil() {
		return ""
	}
	objBytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Sprintf("json Marshal failed: obj=%v, err=%v", obj, err)
	}
	return string(objBytes)
}
