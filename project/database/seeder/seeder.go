package seeder

import (
	"github.com/faizalnurrozi/alterra-agmc/database"
	"gorm.io/gorm"
)

type seed struct {
	DB *gorm.DB
}

func NewSeeder() *seed {
	return &seed{database.GetConnection()}
}

func (s *seed) SeedAll() {
	bookSeeder(s.DB)
	userSeeder(s.DB)
}

func (s *seed) DeleteAll() {
	s.DB.Exec("DELETE FROM books")
	s.DB.Exec("DELETE FROM users")
}
