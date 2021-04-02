package repository

import "github.com/1k-ct/twitter-dem/app/domain/model"

type AccountRepository interface {
	FindByID(ID string) (*model.User, error)
	RegisterUserAccount(*model.User) error
}