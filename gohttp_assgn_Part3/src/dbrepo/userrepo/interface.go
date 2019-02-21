package userrepo

import ("domain"
	"github.com/satori/go.uuid")

//Reader read from db
type Reader interface {
	Get(id string) (*domain.Restaurant, error)
	GetAll() ([]*domain.Restaurant, error)
	//Regex Substring Match on the name field
	FindByName(name string) ([]*domain.Restaurant, error)
}

//Writer  write to db
type Writer interface {
	//Create Or update
	Store(b *domain.Restaurant) (uuid.UUID, error)
	Delete(id string) error
}

//Filter Find objects by additional filters
type Filter interface {
	FindByTypeOfFood(foodType string) ([]*domain.Restaurant, error)
	FindByTypeOfPostCode(postCode string) ([]*domain.Restaurant, error)
	//Search --> across all string fields regex match with case insensitive
	//substring match accross all string fields
	Search(query string) ([]*domain.Restaurant, error)
}

//Repository db interface
type Repository interface {
	Reader
	Writer
	Filter
}

