package model

// EdgexInfo ...
type EdgexInfo struct {
	EdgexID          int64  `json:"edgex_id"`
	UserID           int64  `json:"user_id"`
	UserName         string `json:"username"`
	EdgexName        string `json:"edgex_name"`
	Prefix           string `json:"prefix"`
	Address          string `json:"address"`
	Status           int32  `json:"status"`
	CreatedTimeStamp int64  `json:"created_timestamp"`
	CreatedTime      string `json:"created_time"`
	Description      string `json:"description"`
	Location         string `json:"location"`
	Extra            string `json:"extra"`
	IsFollow         bool   `json:"is_follow"`
}
