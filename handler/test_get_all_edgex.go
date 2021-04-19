package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nju-iot/edgex_admin/dal"
	"github.com/nju-iot/edgex_admin/logs"
	"github.com/nju-iot/edgex_admin/resp"
	"github.com/nju-iot/edgex_admin/wrapper"
)

// GetAllEdgex ...
func GetAllEdgex(c *gin.Context) (out *wrapper.JsonOutput) {
	h := NewGetAllEdgexHandler(c)
	h.Process()
	if len(h.Errors) > 0 {
		return wrapper.SampleJson(c, resp.RESP_CODE_SEVER_EXCEPTION, nil)
	}
	return wrapper.SampleJson(c, resp.RESP_CODE_SUCCESS, h.Resp)
}

type getAllEdgexHandler struct {
	ReqCtx *gin.Context
	Errors []error
	Resp   []*dal.EdgexServiceItem
}

func NewGetAllEdgexHandler(c *gin.Context) *getAllEdgexHandler {
	return &getAllEdgexHandler{
		ReqCtx: c,
	}
}

func (h *getAllEdgexHandler) Process() {
	edgexList, err := dal.GetAllEdgex()
	if err != nil {
		logs.Warn("[getAllEdgexHandler-Process] getAllEdgex faliled: err = %v", err)
		h.Errors = append(h.Errors, err)
		return
	}
	h.Resp = edgexList
}
