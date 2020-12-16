package models

// 权限
type Permission struct {
	BaseModel
	Name string `gorm:"not null;default: '';uniqueIndex" validate:"required"` // 权限名称，必填，唯一
}
