package dal

import (
	"github.com/nju-iot/edgex_admin/caller"
	"github.com/nju-iot/edgex_admin/logs"
)

const (
	TABLE_EDGEX_SERVICE_ITEM = "edgex_service_item"
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

// test
// func InsertTest() error {
// 	item := map[string]interface{}{
// 		"user_id":     123456789,
// 		"edgex_name":  "edgex_1",
// 		"prefix":      "edgex_1",
// 		"description": "测试数据库联通",
// 		"status":      1,
// 		"deleted":     0,
// 		"location":    `{"province":"江苏", "city":"南京市", "longitude":"123", "latitude":"123"}`,
// 	}
// 	dbRes := caller.EdgexDB.Debug().Model(&EdgexServiceItem{}).Create(item)
// 	if dbRes.Error != nil {
// 		logs.Error("[InsertTest] db create row false: err=%v", dbRes.Error)
// 		return dbRes.Error
// 	}
// 	return nil
// }

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
