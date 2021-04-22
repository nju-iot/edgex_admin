package edgex

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nju-iot/edgex_admin/caller"
	"github.com/nju-iot/edgex_admin/dal"
	"github.com/nju-iot/edgex_admin/logs"
	"github.com/nju-iot/edgex_admin/resp"
	"github.com/nju-iot/edgex_admin/wrapper"
)

// RelationEdgexParams ...
type RelationEdgexParams struct {
	UserID    int64  `form:"user_id" json:"user_id" binding:"required"`
	EdgexID   int64  `form:"edgex_id" json:"edgex_id" binding:"required"`
	Username  string `form:"username" json:"username"`
	EdgexName string `form:"edgex_name" json:"edgex_name"`
}

type relationEdgexHandler struct {
	Ctx           *gin.Context
	RelatedEntity *dal.EdgexRelatedUser
	Params        RelationEdgexParams
}

func buildRelationEdgexHandler(c *gin.Context) *relationEdgexHandler {
	return &relationEdgexHandler{
		Ctx: c,
	}
}

// CheckParams ...
func (h *relationEdgexHandler) CheckParams() error {

	err := h.Ctx.Bind(&h.Params)
	if err != nil {
		logs.Error("[relationEdgexHandler-checkParams] params-err: err=%v", err)
		return err
	}
	return nil
}

// Follow ...
func (h *relationEdgexHandler) Follow() (err error) {
	if h.RelatedEntity != nil && h.RelatedEntity.Status == dal.Status_Follow {
		logs.Warn("[relationEdgexHandler-Follow] user followed: user_id=%v, edgex_id=%v", h.Params.UserID, h.Params.EdgexID)
		return
	}
	// Exist
	if h.RelatedEntity != nil {
		filedsMap := map[string]interface{}{"status": dal.Status_Follow}
		err = dal.UpdateEdgexRelatedUser(caller.EdgexDB, h.RelatedEntity.ID, filedsMap)
		if err != nil {
			logs.Error("[relationEdgexHandler-Follow] update follow status failed: filedsMap=%+v, err=%v", filedsMap, err)
			return
		}
		return
	}

	// Create
	entity := &dal.EdgexRelatedUser{
		UserID:       h.Params.UserID,
		Username:     h.Params.Username,
		EdgexID:      h.Params.EdgexID,
		EdgexName:    h.Params.EdgexName,
		Status:       dal.Status_Follow,
		CreatedTime:  time.Now(),
		ModifiedTime: time.Now(),
	}
	err = dal.AddEdgexRelatedUser(caller.EdgexDB, entity)
	if err != nil {
		logs.Error("[relationEdgexHandler-Follow] create follow record: err=%v", err)
		return
	}
	return
}

// UnFollow
func (h *relationEdgexHandler) UnFollow() (err error) {
	if h.RelatedEntity.Status == dal.Status_UnFollow {
		logs.Warn("[relationEdgexHandler-UnFollow] user unfollowed: user_id=%v, edgex_id=%v", h.Params.UserID, h.Params.EdgexID)
		return
	}
	filedsMap := map[string]interface{}{"status": dal.Status_UnFollow}
	err = dal.UpdateEdgexRelatedUser(caller.EdgexDB, h.RelatedEntity.ID, filedsMap)
	if err != nil {
		logs.Error("[relationEdgexHandler-UnFollow] unfollow failed: err=%v", err)
		return
	}
	return
}

// FollowEdgex ...
func FollowEdgex(c *gin.Context) (out *wrapper.JsonOutput) {

	h := buildRelationEdgexHandler(c)

	// Step1. checkParams
	err := h.CheckParams()
	if err != nil {
		logs.Error("[FollowEdgex] params-err: err=%v", err)
		return wrapper.SampleJson(c, resp.RESP_CODE_PARAMS_ERROR, nil)
	}

	// Step2. 获取Follow记录
	h.RelatedEntity, err = dal.FindEdgexRelatedUserByUserIDAndEdgexID(h.Params.EdgexID, h.Params.UserID)
	if err != nil {
		logs.Error("[FollowEdgex] find edgexRelatedUser failed: err=%v", err)
		return wrapper.SampleJson(c, resp.RESP_CODE_DB_ERROR, nil)
	}

	// Step3. Follow
	err = h.Follow()
	if err != nil {
		logs.Error("[FollowEdgex] Follow failed: err=%v", err)
		return wrapper.SampleJson(c, resp.RESP_CODE_DB_ERROR, nil)
	}

	return wrapper.SampleJson(c, resp.RESP_CODE_SUCCESS, "已关注")
}

// UnFollowEdgex ...
func UnFollowEdgex(c *gin.Context) (out *wrapper.JsonOutput) {

	h := buildRelationEdgexHandler(c)

	// Step1. checkParams
	err := h.CheckParams()
	if err != nil {
		logs.Error("[UnFollowEdgex] params-err: err=%v", err)
		return wrapper.SampleJson(c, resp.RESP_CODE_PARAMS_ERROR, nil)
	}

	// Step2. 获取Follow记录
	h.RelatedEntity, err = dal.FindEdgexRelatedUserByUserIDAndEdgexID(h.Params.EdgexID, h.Params.UserID)
	if err != nil {
		logs.Error("[UnFollowEdgex] find edgexRelatedUser failed: err=%v", err)
		return wrapper.SampleJson(c, resp.RESP_CODE_DB_ERROR, nil)
	}
	if h.RelatedEntity == nil {
		logs.Error("[UnFollowEdgex] UnFollow failed: don't have follow record: edgex_id=%v, user_id=%v",
			h.Params.EdgexID, h.Params.UserID)
		return wrapper.SampleJson(c, resp.RESP_CODE_PARAMS_ERROR, nil)
	}

	// Step3. UnFollow
	err = h.UnFollow()
	if err != nil {
		logs.Error("[UnFollowEdgex] UnFollow failed: err=%v", err)
		return wrapper.SampleJson(c, resp.RESP_CODE_DB_ERROR, nil)
	}
	return wrapper.SampleJson(c, resp.RESP_CODE_SUCCESS, "已取消关注")
}
