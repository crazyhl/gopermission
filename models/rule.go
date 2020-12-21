package models

import (
	"github.com/crazyhl/gopermission/config"
	"github.com/crazyhl/gopermission/utils/validator"
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

// 角色添加权限
func (r *Rule) AttachPermission(p []Permission) (*Rule, error) {
	db := config.GetConfig().GormDb

	err := db.Model(r).Association("Permissions").Append(p)

	if err != nil {
		return r, err
	}
	return r, nil
}

// 将角色权限替换为新的权限列表
func (r *Rule) ReplacePermission(p []Permission) (*Rule, error) {
	db := config.GetConfig().GormDb

	err := db.Model(r).Association("Permissions").Replace(p)

	if err != nil {
		return r, err
	}
	return r, nil
}

// 清空角色的所有权限
func (r *Rule) ClearPermission() (*Rule, error) {
	db := config.GetConfig().GormDb

	err := db.Model(r).Association("Permissions").Clear()

	if err != nil {
		return r, err
	}
	return r, nil
}
