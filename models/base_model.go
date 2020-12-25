package models

type BaseModel struct {
	ID        uint32 `gorm:"autoIncrement;primaryKey"`
	CreatedAt uint32 `gorm:"autoCreateTime`      // 使用时间戳秒数填充创建时间
	UpdatedAt uint32 `gorm:"autoUpdateTime`      // 使用时间戳秒数填充创建时间
	Status    uint8  `gorm:"default:1;not null"` // 状态
}
