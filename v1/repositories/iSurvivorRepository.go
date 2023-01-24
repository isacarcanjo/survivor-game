package repositories

import (
	"serve/v1/entitys"
)

// ISurvivorRepository Repository interface
type ISurvivorRepository interface {
	Reader
	Writer
}

type SurvivorDBConn interface {
	Connect() (interface{}, error)
}

//Reader interface
type Reader interface {
	Find(id string) (entitys.SurvivorEntity, error)
	FindAll() ([]entitys.SurvivorEntity, error)
	// Search(query string) ([]*entitys.SurvivorEntity, error)
	// List() ([]*entitys.SurvivorEntity, error)
}

//Writer survivor writer
type Writer interface {
	Create(e entitys.SurvivorEntity) (entitys.SurvivorEntity, error)
	Update(e entitys.SurvivorEntity) (entitys.SurvivorEntity, error)
	// FindUpdate(e entitys.SurvivorEntity) (entitys.SurvivorEntity, error)
	// Delete(id int) error
}
