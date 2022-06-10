package migrations

import "gorm.io/gorm"

type Jwt struct {
	gorm.Model
	ID           int32  `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserID       int32  `gorm:"type:int(12);not null" json:"user_id"`
	Token        string `gorm:"type:varchar(225);not null" json:"token"`
	RefreshToken string `gorm:"type:varchar(225);not null" json:"refresh_token"`
}
