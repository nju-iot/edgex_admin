package dal

import (
	"testing"

	"github.com/nju-iot/edgex_admin/caller"
	"github.com/nju-iot/edgex_admin/config"
)

func init() {
	config.InitConfig()
	caller.InitClient()
}

func TestGetAllEdgex(t *testing.T) {
	edgexList, err := GetAllEdgex()
	t.Logf("err = %v", err)
	for i, edgex := range edgexList {
		t.Logf("index=%v, edgex = %+v", i, edgex)
	}
}
