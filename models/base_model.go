package models

type BaseModel struct {
	ID        uint  `gorm:"autoIncrement;primaryKey"`
	CreatedAt int64 `gorm:"autoCreateTime;not null;default:0"` // 使用时间戳秒数填充创建时间
	UpdatedAt int64 `gorm:"autoCreateTime;not null;default:0"` // 使用时间戳秒数填充创建时间
}
