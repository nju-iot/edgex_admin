package dal

import (
	"github.com/nju-iot/edgex_admin/caller"
	"github.com/nju-iot/edgex_admin/logs"
)

// EdgexServiceItem ...
type EdgexServiceItem struct {
	Id           int64  `gorm:"column:id" json:"id"`
	UserId       int64  `gorm:"column:user_id" json:"user_id"`
	EdgexName    string `gorm:"column:edgex_name" json:"edgex_name"`
	Prefix       string `gorm:"column:prefix" json:"prefix"`
	Status       int32  `gorm:"column:status" json:"status"`
	Deleted      int32  `gorm:"column:deleted" json:"deleted"`
	CreatedTime  string `gorm:"column:created_time" json:"created_time"`
	ModifiedTime string `gorm:"column:modified_time" json:"modified_time"`
	Description  string `gorm:"description:modified_time" json:"description"`
	Location     string `gorm:"column:location" json:"location"`
	Extra        string `gorm:"column:extra" json:"extra"`
}

// GetAllEdgex [test]
func GetAllEdgex() (edgexList []*EdgexServiceItem, err error) {
	edgexList = make([]*EdgexServiceItem, 0)
	dbRes := caller.EdgexDB.Debug().
		Model(&EdgexServiceItem{}).
		Where("deleted = 0").
		Find(&edgexList)
	if dbRes.Error != nil {
		logs.Error("[GetAllEdgex] query all edgex failed: err=%v", dbRes.Error)
		err = dbRes.Error
		return
	}
	return
}
