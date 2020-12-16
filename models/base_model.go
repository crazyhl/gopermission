package models

type BaseModel struct {
	ID        uint  `gorm:"type:int(10) AUTO_INCREMENT;primaryKey"`
	CreatedAt int64 `gorm:"type:int(10);autoCreateTime"` // 使用时间戳秒数填充创建时间
	UpdatedAt int64 `gorm:"type:int(10);autoCreateTime"` // 使用时间戳秒数填充创建时间
}
