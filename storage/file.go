package storage

import (
	"app/config"
)

type Store struct {
	User *userRepo
}

func NewFileJson(cfg *config.Config) (*Store, error) {
	return &Store{
		User: NewUserRepo(cfg.UserFileName),
	}, nil
}
