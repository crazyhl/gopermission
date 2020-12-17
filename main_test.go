package gopermission

import (
	"github.com/crazyhl/gopermission/v1/base_struct"
	"github.com/crazyhl/gopermission/v1/models"
	permissionService "github.com/crazyhl/gopermission/v1/service/permission"
	ruleService "github.com/crazyhl/gopermission/v1/service/rule"
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
	register()
}

func Test_Add_Permission(t *testing.T) {
	register()
	url := "/user/keys/:id/:name"
	p := &models.Permission{
		Name: "TestPermission",
		Url:  url,
	}
	var err error
	p, err = p.Add()

	t.Log(p)
	t.Log(err)
}

func Test_Update_Permission(t *testing.T) {
	register()
	p := permissionService.FindById(1)

	t.Log(p)
	p.Name = "TestPermissionUpdate"
	p.Url = "/user/keys/123/name"

	var err error
	p, err = p.Update()

	t.Log(p)
	t.Log(err)
}

func Test_Delete_Permission(t *testing.T) {
	register()
	p := permissionService.FindById(1)

	count, err := p.Delete()
	t.Log(count)
	t.Log(err)
}

func Test_Add_Rule(t *testing.T) {
	register()
	r := &models.Rule{
		Name: "test",
	}
	var err error
	r, err = r.Add()

	t.Log(r)
	t.Log(err)
}

func Test_Update_Rule(t *testing.T) {
	register()
	r := ruleService.FindById(1)

	t.Log(r)
	r.Name = "test update"

	var err error
	r, err = r.Update()

	t.Log(r)
	t.Log(err)
}

func Test_Delete_Rule(t *testing.T) {
	register()
	r := ruleService.FindById(4)

	count, err := r.Delete()
	t.Log(count)
	t.Log(err)
}

func Test_Add_Rule_With_Permission(t *testing.T) {
	register()
	r := &models.Rule{
		Name: "test",
	}
	permissions := []models.Permission{}
	permissions = append(permissions, models.Permission{
		Name: "testp",
		Url:  "/fdsa/fdsa",
	})
	r.Permissions = permissions

	var err error
	r, err = r.Add()

	t.Log(r)
	t.Log(err)
}

func Test_Add_Rule_With_Exist_Permission(t *testing.T) {
	register()
	r := &models.Rule{
		Name: "test2",
	}
	permissions := []models.Permission{}
	permissions = append(permissions, *permissionService.FindById(2))
	r.Permissions = permissions

	var err error
	r, err = r.Add()

	t.Log(r)
	t.Log(err)
}

func Test_Replace_Rule_Permission(t *testing.T) {
	register()
	r := ruleService.FindById(2)
	permissions := []models.Permission{}
	permissions = append(permissions, models.Permission{
		Name: "testp3",
		Url:  "/fdsa/fdsa",
	})
	r.Permissions = permissions
	t.Log(r.Permissions)
	var err error
	r, err = r.Update()

	t.Log(r)
	t.Log(err)
}

// --------------------- private function -----------------------------

func register() {
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
