package request_test

import (
	"testing"

	"github.com/2981047480/tinybox/common/request"
)

func TestGet(t *testing.T) {
	r := request.NewHttpRequest()
	u, err := r.InitUrl("https://www.baidu.com", "")
	if err != nil {
		t.Fatal(err)
	}
	// u.RawQuery = "page=1&size=999"
	full_url := u.String()

	res, err := r.Get(full_url, "", "")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}
