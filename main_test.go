package gopermission

import (
	"github.com/crazyhl/gopermission/v1/base_struct"
	"testing"
)

func Test_Route_Parse(t *testing.T) {
	path := "/user/keys/:id/:name"
	url := "/user/keys/123/name"
	value := GetParams(path, url, "id")

	t.Log(value)
	t.Log(value == "123")
	value = GetParams(path, url, "name")

	t.Log(value)
	t.Log(value == "name")
}

func Test_Connect_Db(t *testing.T) {
	RegisterWithConfig(base_struct.DbConfig{
		Username: "root",
		Password: "123456",
		Host:     "127.0.0.1",
		Port:     "33060",
		Database: "finance",
		Charset:  "utf8mb4",
		Location: "Asia%2fShanghai",
	}, func(modelName string, getModelFieldName string, paramValue string, condition string) bool {
		return true
	})
}
