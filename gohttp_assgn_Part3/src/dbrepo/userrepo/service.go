package userrepo

import (
	"domain"
	"utils"
	"github.com/satori/go.uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

/*type Reader interface {
	Get(id domain.ID) (*domain.Restaurant, error)
	GetAll() ([]*domain.Restaurant, error)
	//Regex Substring Match on the name field
	FindByName(name string) ([]*domain.Restaurant, error)
}

//Writer  write to db
type Writer interface {
	//Create Or update
	Store(b *domain.Restaurant) (domain.ID, error)
	Delete(id domain.ID) error
}

//Filter Find objects by additional filters
type Filter interface {
	FindByTypeOfFood(foodType string) ([]*domain.Restaurant, error)
	FindByTypeOfPostCode(postCode string) ([]*domain.Restaurant, error)
	//Search --> across all string fields regex match with case insensitive
	//substring match accross all string fields
	Search(query string) ([]*domain.Restaurant, error)
}
*/

func (s *Service) GetAll() ([]*domain.Restaurant, error) {
	return s.repo.GetAll()
}


func (s *Service)FindByName(name string) ([]*domain.Restaurant, error){
	return s.repo.FindByName(name)
}

func (s *Service) Get(ID string) (*domain.Restaurant, error) {
	return s.repo.Get(ID)
}

func (s *Service) Store(b *domain.Restaurant) (uuid.UUID, error) {
	check_id,_ := uuid.FromString("")
	if (b.DBID == check_id){
		b.DBID = utils.NewUUID()	
	}
	return s.repo.Store(b)
	
}


func (s *Service) Delete (id string) error{
		
	return s.repo.Delete(id)
}

func (s * Service)FindByTypeOfFood(foodType string) ([]*domain.Restaurant, error){
	return s.repo.FindByTypeOfFood(foodType)

}

func (s * Service)FindByTypeOfPostCode(postCode string) ([]*domain.Restaurant, error){
	return s.repo.FindByTypeOfPostCode(postCode)
}

func (s * Service)Search(query string) ([]*domain.Restaurant, error){
	return s.repo.Search(query)
}

