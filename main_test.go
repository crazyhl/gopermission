package gopermission

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crazyhl/gopermission/base_struct"
	"github.com/crazyhl/gopermission/config"
	"github.com/crazyhl/gopermission/listener/conditon"
	"github.com/crazyhl/gopermission/models"
	"github.com/crazyhl/gopermission/parser"
	permissionService "github.com/crazyhl/gopermission/service/permission"
	ruleService "github.com/crazyhl/gopermission/service/rule"
	"github.com/fatih/structs"
	"strings"
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
		Name:                "TestPermission",
		Url:                 url,
		Method:              "GEt",
		ModelName:           "TestModel",
		UrlParamName:        "id",
		ModelCheckCondition: "a==b0",
	}
	var err error
	err = p.Add()

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
	err = p.Update()

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
		Name: "test1111",
	}
	var err error
	err = r.Add()

	t.Log(r)
	t.Log(err)
}

func Test_Update_Rule(t *testing.T) {
	register()
	r := ruleService.FindById(19)

	t.Log(r)
	r.Name = "test update"

	var err error
	err = r.Update()

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

func Test_Get_Rule(t *testing.T) {
	register()
	r := ruleService.FindById(11)

	t.Log(r)
	t.Log(StructToMap(r))
}

func Test_Add_Rule_With_Permission(t *testing.T) {
	register()
	r := &models.Rule{
		Name: "test",
	}
	var permissions []models.Permission
	permissions = append(permissions, models.Permission{
		Name: "testp",
		Url:  "/fdsa/fdsa",
	})
	r.Permissions = permissions

	var err error
	err = r.Add()

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
	err = r.Add()

	t.Log(r)
	t.Log(err)
}

func Test_Attach_Rule_Permission(t *testing.T) {
	register()
	r := ruleService.FindById(1)
	ids := []int{1}
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
	//condition := "model.Uid == user.Id"
	condition := "model.Uid == 3"
	modelData := make(map[string]interface{})
	userData := make(map[string]interface{})
	modelData["Uid"] = 3
	userData["Id"] = "456"
	fmt.Println(modelData)
	fmt.Println(userData)
	result := conditon.GetConditionResult(condition, modelData, userData)
	t.Log(result)
}

func Test_Condition_Not_Equal_Check(t *testing.T) {
	condition := "model.Uid != user.Id"
	modelData := make(map[string]interface{})
	userData := make(map[string]interface{})
	modelData["Uid"] = 4561
	userData["Id"] = "456"
	fmt.Println(modelData)
	fmt.Println(userData)
	result := conditon.GetConditionResult(condition, modelData, userData)
	t.Log(result)
}

func Test_Condition_Deep_Equal_Check(t *testing.T) {
	condition := "model.Article.Uid == user.Id"
	modelData := make(map[string]interface{})
	userData := make(map[string]interface{})
	modelData["Uid"] = 123
	articleMap := make(map[string]interface{})
	articleMap["Uid"] = 4561
	modelData["Article"] = articleMap
	userData["Id"] = "456"
	result := conditon.GetConditionResult(condition, modelData, userData)
	t.Log(result)
	t.Log(fmt.Sprint(nil))
	paramStr := "123"
	paramsArr := strings.Split(paramStr, ".")
	t.Log(paramsArr)
}

func Test_Condition_Larger_Check(t *testing.T) {
	condition := "model.Uid > user.Id"
	//condition := "model.Uid > 3"
	modelData := make(map[string]interface{})
	userData := make(map[string]interface{})
	modelData["Uid"] = 4566
	userData["Id"] = "456"
	fmt.Println(modelData)
	fmt.Println(userData)
	result := conditon.GetConditionResult(condition, modelData, userData)
	t.Log(result)
}

func Test_Condition_Larger_Equal_Check(t *testing.T) {
	condition := "model.Uid >= user.Id"
	//condition := "model.Uid > 3"
	modelData := make(map[string]interface{})
	userData := make(map[string]interface{})
	modelData["Uid"] = 45
	userData["Id"] = "456"
	fmt.Println(modelData)
	fmt.Println(userData)
	result := conditon.GetConditionResult(condition, modelData, userData)
	t.Log(result)
}

func Test_Condition_Less_Check(t *testing.T) {
	condition := "model.Uid < user.Id"
	//condition := "model.Uid > 3"
	modelData := make(map[string]interface{})
	userData := make(map[string]interface{})
	modelData["Uid"] = 456
	userData["Id"] = "456"
	fmt.Println(modelData)
	fmt.Println(userData)
	result := conditon.GetConditionResult(condition, modelData, userData)
	t.Log(result)
}

func Test_Condition_Less_Equal_Check(t *testing.T) {
	condition := "model.Uid <= user.Id"
	//condition := "model.Uid > 3"
	modelData := make(map[string]interface{})
	userData := make(map[string]interface{})
	modelData["Uid"] = 4561
	userData["Id"] = "456"
	fmt.Println(modelData)
	fmt.Println(userData)
	result := conditon.GetConditionResult(condition, modelData, userData)
	t.Log(result)
}

func Test_Condition_In_Check(t *testing.T) {
	condition := "model.Name in user.Categories"
	//condition := "model.Uid > 3"
	modelData := make(map[string]interface{})

	modelData["Id"] = 1
	modelData["Name"] = "c3"
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

	fmt.Println(modelData)
	result := conditon.GetConditionResult(condition, modelData, userMap)
	t.Log(result)
}

func Test_Condition_Or_Check(t *testing.T) {
	condition := "model.Name in user.Categories || model.Uid == 3"
	//condition := "model.Uid > 3"
	modelData := make(map[string]interface{})

	modelData["Uid"] = 1
	modelData["Name"] = "c1"
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

	fmt.Println(modelData)
	result := conditon.GetConditionResult(condition, modelData, userMap)
	t.Log(result)
}

func Test_Condition_And_Check(t *testing.T) {
	condition := "model.Name in user.Categories && model.Uid == 2"
	//condition := "model.Uid > 3"
	modelData := make(map[string]interface{})

	modelData["Uid"] = 1
	modelData["Name"] = "c1"
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
	userMap := StructToMap(user)
	fmt.Println(userMap)

	fmt.Println(modelData)
	result := conditon.GetConditionResult(condition, modelData, userMap)
	t.Log(result)
}

func Test_Condition_Mix_Check(t *testing.T) {
	condition := "(model.Name in user.Categories && model.Uid == 1) || model.Id == 15"
	//condition := "model.Uid > 3"
	modelData := make(map[string]interface{})

	modelData["Uid"] = 1
	modelData["Id"] = 5
	modelData["Name"] = "c111"
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

	fmt.Println(modelData)
	result := conditon.GetConditionResult(condition, modelData, userMap)
	t.Log(result)
}

func Test_Has_Permission(t *testing.T) {
	register()
	gormDb := config.GetConfig().GormDb
	permissions := []models.Permission{}
	gormDb.Find(&permissions)
	userData := make(map[string]interface{})
	userData["Uid"] = 123
	result := HasPermission(userData, "Delete", "/user/keys/:id/:name", "/user/keys/123/name", permissions)
	t.Log(result)
}

// --------------------- private function -----------------------------

func register() {
	RegisterWithConfigAndAutoMigrate(base_struct.DbConfig{
		Username: "root",
		Password: "123456",
		Host:     "127.0.0.1",
		Port:     "33060",
		Database: "finance",
		Charset:  "utf8mb4",
		Location: "Asia%2fShanghai",
	}, func(modelName string, paramValue string, condition string, user map[string]interface{}) bool {
		switch modelName {
		case "Model":
			modelMap := make(map[string]interface{})
			if paramValue == "123" {
				modelMap["Uid"] = 123
				return conditon.GetConditionResult(condition, modelMap, user)
			}

		}

		return false
	})
}
