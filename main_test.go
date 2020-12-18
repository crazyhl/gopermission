package gopermission

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crazyhl/gopermission/v1/base_struct"
	"github.com/crazyhl/gopermission/v1/listener/conditon"
	"github.com/crazyhl/gopermission/v1/models"
	"github.com/crazyhl/gopermission/v1/parser"
	permissionService "github.com/crazyhl/gopermission/v1/service/permission"
	ruleService "github.com/crazyhl/gopermission/v1/service/rule"
	"github.com/fatih/structs"
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

func Test_Attach_Rule_Permission(t *testing.T) {
	register()
	r := ruleService.FindById(6)
	ids := []int{2, 3}
	permissions := permissionService.FindByIds(ids)
	t.Log(permissions)

	var err error
	r, err = r.AttachPermission(permissions)

	t.Log(r)
	t.Log(err)
}

func Test_Replace_Rule_Permission(t *testing.T) {
	register()
	r := ruleService.FindById(6)
	ids := []int{4}
	permissions := permissionService.FindByIds(ids)
	t.Log(permissions)

	var err error
	r, err = r.ReplacePermission(permissions)

	t.Log(r)
	t.Log(err)
}

func Test_Clear_Rule_Permission(t *testing.T) {
	register()
	r := ruleService.FindById(6)

	var err error
	r, err = r.ClearPermission()

	t.Log(r)
	t.Log(err)
}

func Test_Parser(t *testing.T) {
	is := antlr.NewInputStream("(model.uid == user.uid || (model.a <= user.c && model.d in user.abc) )  || model.uid > 3")

	// Create the Lexer
	lexer := parser.NewConditionLexer(is)

	// Read all tokens
	//for {
	//	t := lexer.NextToken()
	//	if t.GetTokenType() == antlr.TokenEOF {
	//		break
	//	}
	//	fmt.Printf("%s (%q)\n",
	//		lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	//}
	steam := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewConditionParser(steam)
	listener := &conditon.ConditionListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Start())
}

type User struct {
	Username   string
	Categories []Cateogory
}

type Cateogory struct {
	Name string
	Id   int
}

func Test_Struct_To_Map(t *testing.T) {
	categories := make([]Cateogory, 0)
	categories = append(categories, Cateogory{
		Name: "c1",
		Id:   1,
	})
	categories = append(categories, Cateogory{
		Name: "c2",
		Id:   2,
	})
	user := User{
		Username:   "aaa",
		Categories: categories,
	}
	userMap := structs.Map(user)
	fmt.Println(userMap)
}

func Test_Condition_Equal_Check(t *testing.T) {

	condition := "model.Uid == user.Id"
	modelData := make(map[interface{}]interface{})
	userData := make(map[interface{}]interface{})
	modelData["Uid"] = 456
	userData["Id"] = "456"
	fmt.Println(modelData)
	fmt.Println(userData)
	result := conditon.GetConditionResult(condition, modelData, userData)
	t.Log(result)
}

func Test_Condition_Not_Equal_Check(t *testing.T) {
	condition := "model.Uid != user.Id"
	modelData := make(map[interface{}]interface{})
	userData := make(map[interface{}]interface{})
	modelData["Uid"] = 123
	userData["Id"] = "456"
	fmt.Println(modelData)
	fmt.Println(userData)
	result := conditon.GetConditionResult(condition, modelData, userData)
	t.Log(result)
}

func Test_Condition_Deep_Equal_Check(t *testing.T) {
	condition := "model.Article.Uid == user.Id"
	modelData := make(map[interface{}]interface{})
	userData := make(map[interface{}]interface{})
	modelData["Uid"] = 123
	articleMap := make(map[interface{}]interface{})
	articleMap["Uid"] = 456
	modelData["Article"] = articleMap
	userData["Id"] = "456"
	fmt.Println(modelData)
	fmt.Println(userData)
	result := conditon.GetConditionResult(condition, modelData, userData)
	t.Log(result)
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
	}, func(modelName string, getModelFieldName string, paramValue string, condition string, user map[interface{}]interface{}) bool {
		return true
	})
}
