package migrations

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	ID       int32   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name     string  `gorm:"type:varchar(100);not null" json:"name"`
	Email    string  `gorm:"type:varchar(100);not null" json:"email"`
	Password string  `gorm:"type:varchar(100);not null" json:"password"`
	Phone    string  `gorm:"type:varchar(100);not null" json:"phone"`
	Address  string  `gorm:"type:varchar(225);not null" json:"address"`
	Avatar   Avatars `gorm:"foreignkey:UserID" json:"avatar"`
	Jwt      Jwt     `gorm:"foreignkey:UserID" json:"jwt"`
}
