package gopermission

import (
	"crypto/md5"
	"fmt"
	"github.com/crazyhl/gopermission/base_struct"
	"github.com/crazyhl/gopermission/bootstrap"
	"github.com/crazyhl/gopermission/config"
	"github.com/crazyhl/gopermission/models"
	"github.com/crazyhl/gopermission/route"
	"github.com/fatih/structs"
	"gorm.io/gorm"
	"strings"
)

var customCheckFunction base_struct.CustomModelCheck

// 注册使用权限系统
func RegisterWithConfigAndAutoMigrate(dbConfig base_struct.DbConfig, checkFunction base_struct.CustomModelCheck) {
	// 初始化数据库
	bootstrap.Db(dbConfig)
	customCheckFunction = checkFunction
}

// 注册使用权限系统
func RegisterWithDbAndAutoMigrate(gormDb *gorm.DB, checkFunction base_struct.CustomModelCheck) {
	// 初始化数据库
	bootstrap.AutoMigrateAndSetToContainer(gormDb)
	customCheckFunction = checkFunction
}

func RegisterWithDb(gormDb *gorm.DB, checkFunction base_struct.CustomModelCheck) {
	bootstrap.SetToContainer(gormDb)
	customCheckFunction = checkFunction
}

func AutoMigrate(gormDb *gorm.DB) {
	bootstrap.AutoMigrate(gormDb)
}

// 获取权限
func GetPermission(path string, requestMethod *string) []models.Permission {
	// 参数 path 是 传入 权限的 url 字段
	pathBytes := []byte(path)
	md5Bytes := md5.Sum(pathBytes)
	md5Str := fmt.Sprintf("%x", md5Bytes)
	db := config.GetConfig().GormDb

	permissions := []models.Permission{}

	requestMethodWhere := []string{"*", ""}
	if requestMethod != nil {
		requestMethodWhere = []string{strings.ToLower(*requestMethod), "*", ""}
	}

	db.Where("url_md5 = ?", md5Str).Where("method IN ?", requestMethodWhere).Find(&permissions)

	return permissions
}

// 权限检测
func HasPermission(user map[string]interface{}, requestMethod string, path string, uri string, userPermissions []models.Permission) bool {
	bindPermissions := GetPermission(path, &requestMethod)
	// 如果 路径没有绑定任何权限，那么就直接通过
	if len(bindPermissions) == 0 {
		return true
	}
	// 如果获取到了权限就进行匹配
	for _, bindPermission := range bindPermissions {
		for _, userPermission := range userPermissions {
			if bindPermission.Name == userPermission.Name && (bindPermission.Method == "" || bindPermission.Method == "*" || strings.ToLower(requestMethod) == bindPermission.Method) {
				// 如果用户包含这个权限就进行后续判定
				// 如果这个权限没有绑定模型就可以直接通过了
				if bindPermission.ModelName == "" {
					return true
				}
				// 如果绑定了权限，那么就要进行后续判定了
				paramValue := GetParams(path, uri, bindPermission.UrlParamName)
				customCheckResult := customCheckFunction(bindPermission.ModelName, paramValue, bindPermission.ModelCheckCondition, user)
				// 只要有一个条件返回 true 就可以通过了
				if customCheckResult == true {
					return true
				}
			}
		}
	}
	return false
}

func GetParams(path string, url string, paramKey string) string {
	//var params [route.MaxParams]string
	var ctxParams [route.MaxParams]string
	rp := route.ParseRoute(path)

	rp.GetMatch(url, url, &ctxParams, false)
	return route.Params(rp.Params, ctxParams, paramKey)
}

// 把 struct 转换为 map
func StructToMap(model interface{}) map[string]interface{} {
	return structs.Map(model)
}
