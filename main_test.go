package gopermission

import (
	"testing"
)

func Test_Route_Parse(t *testing.T) {
	path := "/user/keys/:id/:name"
	url := "/user/keys/123/456"
	value := GetParams(path, url, "id")

	t.Log(value)
}
