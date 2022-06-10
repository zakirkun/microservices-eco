package migrations

import "gorm.io/gorm"

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
		&Avatars{},
		&Users{},
		&Jwt{},
	)

	m.db.Migrator().CreateConstraint(&Users{}, "Jwt")
	m.db.Migrator().CreateConstraint(&Users{}, "Avatars")

	return err
}
