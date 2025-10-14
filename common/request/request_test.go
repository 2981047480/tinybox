package request_test

import (
	"testing"

	"github.com/2981047480/tinybox/common/request"
)

func TestGet(t *testing.T) {
	r := request.NewHttpRequest()
	r.InitUrl("https://linkcloud-admin.qiniu.io", "/api/proxy/jarvis/v2/nodes")
	res, err := r.Get()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
