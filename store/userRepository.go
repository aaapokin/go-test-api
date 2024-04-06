package store

import (
	"app/internal/app/models"
	"fmt"
)

type UserRepository struct {
	store *Store
}

var tableUser string = "users"

func (repository *UserRepository) Create(model *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (id,login, password) VALUES ($1, $2, $3) RETURNING id", tableUser)
	if err := repository.store.DB.QueryRow(
		query,
		model.Login,
		model.Password,
	).Scan(&model.ID); err != nil {
		return nil, err
	}
	return model, nil
}

func (repository *UserRepository) FindById(id string) (*models.User, bool, error) {
	model := &models.User{}
	ok := false
	query := fmt.Sprintf("SELECT id, login FROM %s where id=$1", tableUser)
	if err := repository.store.DB.QueryRow(
		query,
		id,
	).Scan(&model.ID, &model.Login); err != nil {
		return nil, ok, err
	}
	if model.ID != "" {
		ok = true
	}
	return model, ok, nil
}

func (repository *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	model := &models.User{}
	ok := false
	query := fmt.Sprintf("SELECT id, login, password FROM %s where login=$1", tableUser)
	if err := repository.store.DB.QueryRow(
		query,
		login,
	).Scan(&model.ID, &model.Login, &model.Password); err != nil {
		return nil, ok, err
	}
	if model.ID != "" {
		ok = true
	}
	return model, ok, nil
}
