package implementations

import (
	"encoding/json"
	"fmt"
	"serve/v1/entitys"

	scribble "github.com/nanobox-io/golang-scribble"
)

type SimDBProvider struct {
	SurvivorCollection string
}

func (s SimDBProvider) connect() (*scribble.Driver, error) {
	db, errDb := scribble.New("./", nil)
	if errDb != nil {
		fmt.Println("Error", errDb)
		return nil, errDb
	}
	return db, nil
}

func (u SimDBProvider) Create(e entitys.SurvivorEntity) (entitys.SurvivorEntity, error) {
	db, errDb := u.connect()
	if errDb != nil {
		fmt.Println("Error", errDb)
	}

	if err := db.Write(u.SurvivorCollection, e.Id, e); err != nil {
		fmt.Println("Error", err)
		return entitys.SurvivorEntity{}, nil
	}
	return e, nil
}

func (u SimDBProvider) Find(id string) (entitys.SurvivorEntity, error) {
	db, errDb := u.connect()
	if errDb != nil {
		fmt.Println("Error", errDb)
	}
	survivor := entitys.SurvivorEntity{}
	if err := db.Read(u.SurvivorCollection, id, &survivor); err != nil {
		fmt.Println("Error", err)
		return survivor, entitys.ErrNotFound
	}
	return survivor, nil
}

func (u SimDBProvider) Update(e entitys.SurvivorEntity) (entitys.SurvivorEntity, error) {
	db, errDb := u.connect()
	if errDb != nil {
		return entitys.SurvivorEntity{}, nil
	}

	if errDb != nil {
		fmt.Println("Error", errDb)
	}
	if err := db.Write(u.SurvivorCollection, e.Id, e); err != nil {
		fmt.Println("Error", err)
		return entitys.SurvivorEntity{}, nil
	}
	return e, nil
}

func (u SimDBProvider) FindAll() ([]entitys.SurvivorEntity, error) {
	db, errDb := u.connect()
	var survivors []entitys.SurvivorEntity
	if errDb != nil {
		fmt.Println("Error", errDb)
		return survivors, nil
	}
	records, err := db.ReadAll(u.SurvivorCollection)
	if err != nil {
		fmt.Println("Error", err)
		return survivors, nil
	}
	for _, f := range records {
		survivorFound := entitys.SurvivorEntity{}
		if errUnmarshal := json.Unmarshal([]byte(f), &survivorFound); errUnmarshal != nil {
			fmt.Println("Error", errUnmarshal)
			return survivors, errUnmarshal
		}
		survivors = append(survivors, survivorFound)
	}
	return survivors, nil
}
