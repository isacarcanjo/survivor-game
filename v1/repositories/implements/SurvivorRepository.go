package repositories

import (
	"serve/v1/entitys"
	"serve/v1/repositories"
)

type SurvivorRepository struct {
	db repositories.ISurvivorRepository
}

//NewSurvivorRepository create new repository
func NewSurvivorRepository(db repositories.ISurvivorRepository) *SurvivorRepository {
	return &SurvivorRepository{db}
}

//Create survivor
func (r *SurvivorRepository) Create(e entitys.SurvivorEntity) (entitys.SurvivorEntity, error) {
	survivor, err := r.db.Create(e)
	if err != nil {
		return entitys.SurvivorEntity{}, err
	}
	return survivor, nil
}

//Find survivor
func (r *SurvivorRepository) Find(id string) (entitys.SurvivorEntity, error) {
	survivor, err := r.db.Find(id)
	if err != nil {
		return entitys.SurvivorEntity{}, err
	}
	return survivor, nil
}

//Update survivor
func (r *SurvivorRepository) Update(e entitys.SurvivorEntity) (entitys.SurvivorEntity, error) {
	survivor, err := r.db.Update(e)
	if err != nil {
		return entitys.SurvivorEntity{}, err
	}
	return survivor, nil
}

//Update survivor
// func (r *SurvivorRepository) FindUpdate(e entitys.SurvivorEntity) (entitys.SurvivorEntity, error) {
// 	survivor, err := r.db.FindUpdate(e)
// 	if err != nil {
// 		return entitys.SurvivorEntity{}, err
// 	}
// 	return survivor, nil
// }

//FindAll survivors
func (r *SurvivorRepository) FindAll() ([]entitys.SurvivorEntity, error) {
	var survivors []entitys.SurvivorEntity
	survivors, err := r.db.FindAll()
	if err != nil {
		return survivors, err
	}
	return survivors, nil
}
