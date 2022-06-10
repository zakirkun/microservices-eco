package entity

import "gorm.io/gorm"

type Avatars struct {
	gorm.Model
	ID     int32  `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserID int32  `gorm:"type:int(12);not null" json:"user_id"`
	Avatar string `gorm:"type:varchar(225);not null" json:"avatar"`
}
