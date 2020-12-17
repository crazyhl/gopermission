package models

import (
	"github.com/crazyhl/gopermission/v1/config"
	"github.com/crazyhl/gopermission/v1/utils/validator"
	"gorm.io/gorm/clause"
)

// 角色
type Rule struct {
	BaseModel
	Name        string       `gorm:"not null;default: '';uniqueIndex" validate:"required"` // 权限名称，必填，唯一
	Permissions []Permission `gorm:"many2many:rule_permissions;"`
}

// 增加
func (r *Rule) Add() (*Rule, error) {
	errs := validator.Validate(r)
	if len(errs) > 0 {
		return nil, errs[0]
	}

	db := config.GetConfig().GormDb
	result := db.Create(r)
	if result.Error != nil {
		return nil, result.Error
	}

	return r, nil
}

// 更新 - 这个只更新本体，不要更新关联权限相关，权限相关用后面权限相关的方法
func (r *Rule) Update() (*Rule, error) {
	errs := validator.Validate(r)
	if len(errs) > 0 {
		return nil, errs[0]
	}

	db := config.GetConfig().GormDb
	result := db.Save(r)
	if result.Error != nil {
		return nil, result.Error
	}

	return r, nil
}

// 删除
func (r *Rule) Delete() (int64, error) {
	db := config.GetConfig().GormDb
	result := db.Select(clause.Associations).Delete(r)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
