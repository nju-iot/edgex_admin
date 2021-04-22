package edgex

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nju-iot/edgex_admin/constdef"
	"github.com/nju-iot/edgex_admin/dal"
	"github.com/nju-iot/edgex_admin/logs"
	"github.com/nju-iot/edgex_admin/model"
	"github.com/nju-iot/edgex_admin/resp"
	"github.com/nju-iot/edgex_admin/utils"
	"github.com/nju-iot/edgex_admin/wrapper"
)

const (
	Action_All    = "all"
	Action_Me     = "me"
	Action_Follow = "follow"
)

// SearchEdgexParams ...
type SearchEdgexParams struct {
	Action  string `form:"action" json:"action"`
	UserID  int64  `form:"user_id" json:"user_id"`
	Keyword string `form:"keyword" json:"keyword"`
	Status  int    `form:"status" json:"status"`
	Offset  int    `form:"offset" json:"offset"`
	Count   int    `form:"count" json:"count"`
}

type searchEdgexHandler struct {
	Ctx       *gin.Context
	Params    SearchEdgexParams
	EdgexList []*model.EdgexInfo
}

func buildSearchEdgexHandler(c *gin.Context) *searchEdgexHandler {
	return &searchEdgexHandler{
		Ctx:       c,
		EdgexList: make([]*model.EdgexInfo, 0),
	}
}

// SearchEdgex ...
func SearchEdgex(c *gin.Context) (out *wrapper.JsonOutput) {

	h := buildSearchEdgexHandler(c)

	// Step1. checkParams
	err := h.CheckParams()
	if err != nil {
		logs.Error("[SearchEdgex] params-err: err=%v", err)
		return wrapper.SampleJson(c, resp.RESP_CODE_PARAMS_ERROR, nil)
	}

	// Step2. search
	err = h.Process()
	if err != nil {
		logs.Error("[SearchEdgex] params-err: err=%v", err)
		return wrapper.SampleJson(c, resp.RESP_CODE_DB_ERROR, nil)
	}

	return wrapper.SampleJson(c, resp.RESP_CODE_SUCCESS, h.EdgexList)
}

func (h *searchEdgexHandler) CheckParams() error {

	err := h.Ctx.Bind(&h.Params)
	if err != nil {
		logs.Error("[searchEdgexHandler-checkParams] params-err: err=%v", err)
		return err
	}

	if h.Params.Action == "" {
		h.Params.Action = Action_All
	}

	if !utils.InStringSlice(h.Params.Action, []string{Action_All, Action_Me, Action_Follow}) {
		return fmt.Errorf("action is invalid: action=%s", h.Params.Action)
	}

	if (h.Params.Action == Action_Me || h.Params.Action == Action_Follow) && h.Params.UserID == 0 {
		return fmt.Errorf("params error: action=%s but user_id=0", h.Params.Action)
	}

	if h.Params.Count == 0 {
		h.Params.Count = 10
	}
	return nil
}

func (h *searchEdgexHandler) Process() (err error) {

	var (
		edgexIDs []int64
		userIDs  []int64
		keyword  = h.Params.Keyword
	)

	switch h.Params.Action {
	case Action_Me:
		userIDs = []int64{h.Params.UserID}
	case Action_Follow:
		edgexIDs, err = dal.GetFollowEdgexIDs(h.Params.UserID)
		if err != nil {
			return
		}
		if len(edgexIDs) == 0 {
			return nil
		}
	}

	// 指定edgex_id搜索
	if edgexID, err := strconv.ParseInt(keyword, 10, 64); err == nil && edgexID > 0 {
		if h.Params.Action == Action_Follow && !utils.InInt64Slice(edgexID, edgexIDs) {
			return nil
		}
		edgexIDs = []int64{edgexID}
		keyword = ""
	}

	edgexList, err := dal.GetEdgexList(edgexIDs, userIDs, keyword, h.Params.Status, h.Params.Offset, h.Params.Count)

	if err != nil {
		logs.Error("[searchEdgexHandler-Process] GetEdgexList failed: err=%v", err)
		return err
	}

	h.Pack(edgexList)
	return
}

func (h *searchEdgexHandler) Pack(edgexList []*dal.EdgexServiceItem) {

	h.EdgexList = make([]*model.EdgexInfo, 0)

	for _, item := range edgexList {
		h.EdgexList = append(h.EdgexList, &model.EdgexInfo{
			EdgexID:          item.ID,
			EdgexName:        item.EdgexName,
			UserID:           item.UserID,
			UserName:         "徐志乐", //  TODO：@许月洋
			Prefix:           item.Prefix,
			Address:          item.Address,
			Status:           item.Status,
			CreatedTimeStamp: item.CreatedTime.Unix(),
			CreatedTime:      item.CreatedTime.Format(constdef.TimeFormat),
			Description:      item.Description,
			Location:         item.Location,
			Extra:            item.Extra,
		})
	}
}
