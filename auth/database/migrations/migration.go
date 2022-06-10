package migrations

import (
	"github.com/zakirkun/microservices-eco/auth/database/entity"
	"gorm.io/gorm"
)

type Migration struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Migration {
	return &Migration{
		db: db,
	}
}

func (m Migration) Seeder() error {
	err := m.db.Migrator().CreateTable(
		&entity.Avatars{},
		&entity.Users{},
		&entity.Jwt{},
	)

	m.db.Migrator().CreateConstraint(&entity.Users{}, "Jwt")
	m.db.Migrator().CreateConstraint(&entity.Users{}, "Avatars")

	return err
}
