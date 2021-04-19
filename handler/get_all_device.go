package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/nju-iot/edgex_admin/logs"
	"github.com/nju-iot/edgex_admin/resp"
	"github.com/nju-iot/edgex_admin/wrapper"
	"github.com/parnurzeal/gorequest"
)

func GetAllDevice(c *gin.Context) (out *wrapper.JsonOutput) {
	h := NewGetAllDeviceHandler(c)
	h.Process()
	if len(h.Errors) > 0 {
		return wrapper.SampleJson(c, resp.RESP_CODE_SEVER_EXCEPTION, nil)
	}
	return wrapper.SampleJson(c, resp.RESP_CODE_SUCCESS, h.Resp)
}

type getAllDeviceHandler struct {
	ReqCtx *gin.Context
	Errors []error
	Resp   interface{}
}

func NewGetAllDeviceHandler(c *gin.Context) *getAllDeviceHandler {
	return &getAllDeviceHandler{
		ReqCtx: c,
	}
}

func (h *getAllDeviceHandler) Process() {
	request := gorequest.New()
	deviceURL := "http://47.102.192.194:48082/api/v1/device"

	_, bodyBytes, errs := request.Get(deviceURL).EndBytes()
	if len(errs) > 0 {
		logs.Error("[GetAllDevice] request Get failed: errs=%v")
		h.Errors = errs
		return
	}

	data := make([]interface{}, 0)
	err := json.Unmarshal(bodyBytes, &data)
	if err != nil {
		logs.Error("[GetAllDevice] json unmarshal failed: err=%v", err)
		h.Errors = append(h.Errors, err)
		return
	}
	h.Resp = data
}
