package gopermission

import (
	"github.com/crazyhl/gopermission/route"
	"testing"
)

func Test_Route_Parse(t *testing.T) {
	var ctxParams [route.MaxParams]string
	path := "/user/keys/:id/:name"
	rp := GetParams(path)
	t.Log(rp.Params)
	url := "/user/keys/123/456"
	result := rp.GetMatch(url, url, &ctxParams, false)
	t.Log(result)
	t.Log(ctxParams)
	value := route.Params(rp.Params, ctxParams, "id")
	t.Log(value)
	value = route.Params(rp.Params, ctxParams, "name")
	t.Log(value)
}
