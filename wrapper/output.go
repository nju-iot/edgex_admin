package wrapper

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/nju-iot/edgex_admin/resp"
)

type JsonOutput struct {
	context    *gin.Context
	HttpStatus int
	Resp       interface{}
}

func NewJsonOutput(c *gin.Context, httpStatus int, rsp interface{}) *JsonOutput {
	return &JsonOutput{
		context:    c,
		HttpStatus: httpStatus,
		Resp:       rsp,
	}
}

func SampleJson(c *gin.Context, p resp.ErrorCode, data interface{}) *JsonOutput {
	return NewJsonOutput(c, http.StatusOK, resp.NewStdResponse(p, data))
}

func (s *JsonOutput) GetRespRawData() []byte {
	vi := reflect.ValueOf(s.Resp)
	if vi.Kind() == reflect.Ptr && vi.IsNil() {
		return []byte("")
	}
	rawData, _ := json.Marshal(s.Resp)
	return rawData
}

func (s *JsonOutput) Write() {
	s.context.Writer.WriteHeader(s.HttpStatus)
	s.context.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	s.context.Writer.Write(s.GetRespRawData())
}
