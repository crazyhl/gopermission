package gopermission

import (
	"crypto/md5"
	"github.com/crazyhl/gopermission/base_struct"
	"github.com/crazyhl/gopermission/bootstrap"
	"github.com/crazyhl/gopermission/config"
	"github.com/crazyhl/gopermission/models"
	"github.com/crazyhl/gopermission/route"
)

var customCheckFunction base_struct.CustomModelCheck

// 注册使用权限系统
func Register(dbConfig base_struct.DbConfig) {
	// 初始化数据库
	bootstrap.Db(dbConfig)
}

// 获取权限
func GetPermission(path string) []models.Permission {
	// 参数 path 是 传入 权限的 url 字段
	pathBytes := []byte(path)
	md5Bytes := md5.Sum(pathBytes)
	md5Str := string(md5Bytes[:])
	db := config.GetConfig().GormDb

	permissions := []models.Permission{}

	db.Where("url_md5 = ?", md5Str).Find(&permissions)

	return permissions
}

// 权限检测
func HasPermission(path string, userPermissions []models.Permission) bool {
	bindPermissions := GetPermission(path)
	// 如果 路径没有绑定任何权限，那么就直接通过
	if len(bindPermissions) == 0 {
		return true
	}
	// 如果获取到了权限就进行匹配
	for _, bindPermission := range bindPermissions {
		for _, userPermission := range userPermissions {
			if bindPermission.Name == userPermission.Name {
				// 如果用户包含这个权限就进行后续判定
				// 如果这个权限没有绑定模型就可以直接通过了
				if bindPermission.ModelName == "" {
					return true
				}
				// 如果绑定了权限，那么就要进行后续判定了

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
