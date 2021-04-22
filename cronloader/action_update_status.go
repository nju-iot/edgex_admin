package cronloader

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/nju-iot/edgex_admin/caller"
	"github.com/nju-iot/edgex_admin/dal"
	"github.com/nju-iot/edgex_admin/logs"
	"github.com/parnurzeal/gorequest"
)

func updateEdgexStatus() {
	var (
		offset             = 0
		limit              = 30
		allowDBFailedTimes = 5
		curDBFailedTime    = 0
	)

	for curDBFailedTime < allowDBFailedTimes {
		var err error

		// Step1. 扫描edgex服务
		edgexList, err := dal.GetEdgexList([]int64{}, []int64{}, "", 0, offset, limit)
		if err != nil {
			curDBFailedTime++
			logs.Error("[updateEdgexStatus] select from db failed: curDBFailedTimes=%v, err=%v", curDBFailedTime, err)
			continue
		}

		// Step2. 组装pingURL
		pingURLMap := make(map[int64]string)
		for _, edgex := range edgexList {
			if edgex == nil {
				continue
			}
			if edgex.Address == "" || edgex.Prefix == "" {
				logs.Warn("[updateEdgexStatus] edgex service address is invalid: edgex_id=%v, address=%v, prefix=%v",
					edgex.ID, edgex.Address, edgex.Prefix)
				continue
			}
			pingURL := fmt.Sprintf("http://%s/%s/ping", edgex.Address, edgex.Prefix)
			if _, err := url.Parse(pingURL); err != nil {
				logs.Warn("[updateEdgexStatus] pingURL is invalid: edgexID=%v, pingURL=%v", edgex.ID, pingURL)
				continue
			}
			pingURLMap[edgex.ID] = pingURL
		}

		// Step3. http ping
		statusMap := make(map[int64]int32)
		request := gorequest.New()
		for edgexID, pingURL := range pingURLMap {
			resp, body, errs := request.Get(pingURL).End()
			if len(errs) != 0 || resp == nil || resp.StatusCode != http.StatusOK {
				continue
			}
			if body == "pong" {
				statusMap[edgexID] = 1
			}
		}

		// Step4. update服务状态
		err = func() error {

			var updateErr error
			db := caller.EdgexDB.Begin()
			defer func() {
				if updateErr == nil {
					db.Commit()
				} else {
					db.Callback()
				}
			}()
			for _, edgex := range edgexList {
				if edgex == nil || edgex.Status == statusMap[edgex.ID] {
					continue
				}
				var filedsMap = map[string]interface{}{"status": statusMap[edgex.ID]}
				updateErr = dal.UpdateEdgex(db, edgex.ID, filedsMap)
				if updateErr != nil {
					logs.Error("[updateEdgexStatus] update edgex status failed: edgex_id=%v, status=%v, err=%v", edgex.ID, statusMap[edgex.ID])
					return updateErr
				}
			}
			return nil
		}()
		if err != nil {
			logs.Warn("[updateEdgexStatus] update status failed: err=%v", err)
			continue
		}

		logs.Info("[updateEdgexStatus] success: offset=%v, limit=%v", offset, limit)

		// Step5. finished
		if len(edgexList) < limit {
			break
		}
		offset += limit
		time.Sleep(100 * time.Millisecond)
	}
	logs.Info("[updateEdgexStatus] Finished")
}
