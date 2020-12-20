package models

import (
	"github.com/crazyhl/gopermission/v1/config"
	"github.com/crazyhl/gopermission/v1/utils"
	"github.com/crazyhl/gopermission/v1/utils/validator"
	"gorm.io/gorm"
)

// 权限
type Permission struct {
	BaseModel
	Name      string `gorm:"not null;default: '';uniqueIndex" validate:"required"` // 权限名称，必填，唯一
	Url       string `gorm:"not null;default: '';" validate:"required"`            // 绑定的 url，必填，在我看来一个权限如果不绑定url那是无意义的
	UrlMd5    string `gorm:"type:char(32);not null;default: '';index"`             // 绑定 url 的 md5，用来检索
	ModelName string `gorm:"not null;default: '';"`                                // 绑定的模型名称，这个是用来进行进一步验证用的
	//GetModelFieldName   string  `gorm:"not null;default: '';"`                                        // 获取模型用的字段名称，跟绑定模型合用,所以 path 的变量名和 数据库的字段名要一致
	UrlParamName        string  `gorm:"not null;default: '';"` // 从 url 获取参数的字段名
	ModelCheckCondition string  `gorm:"not null;default: '';"` // 绑定的模型名称跟登录用户的判定条件，这个准备实现一个自己的语义逻辑
	Rules               []*Rule `gorm:"many2many:rule_permissions;"`
}

// 增加
func (p *Permission) Add() (*Permission, error) {
	errs := validator.Validate(p)
	if len(errs) > 0 {
		return nil, errs[0]
	}

	db := config.GetConfig().GormDb
	result := db.Create(p)
	if result.Error != nil {
		return nil, result.Error
	}

	return p, nil
}

// 更新
func (p *Permission) Update() (*Permission, error) {
	errs := validator.Validate(p)
	if len(errs) > 0 {
		return nil, errs[0]
	}

	db := config.GetConfig().GormDb
	result := db.Save(p)
	if result.Error != nil {
		return nil, result.Error
	}

	return p, nil
}

// 删除
func (p *Permission) Delete() (int64, error) {
	db := config.GetConfig().GormDb
	result := db.Delete(p)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func (p *Permission) BeforeCreate(_ *gorm.DB) (err error) {
	p.UrlMd5 = utils.GetStrMd5(p.Url)
	return
}
