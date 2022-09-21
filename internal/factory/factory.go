package factory

import (
	"github.com/faizalnurrozi/alterra-agmc/database"
	"github.com/faizalnurrozi/alterra-agmc/internal/repository"
)

type Factory struct {
	UserRepository repository.User
	BookRepository repository.Book
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		repository.NewUserRepository(db),
		repository.NewBookRepository(db),
	}
}
