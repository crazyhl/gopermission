package models

// 角色
type Rule struct {
	BaseModel
	Name        string       `gorm:"not null;default: '';uniqueIndex" validate:"required"` // 权限名称，必填，唯一
	Permissions []Permission `gorm:"many2many:rule_permissions;"`
}
