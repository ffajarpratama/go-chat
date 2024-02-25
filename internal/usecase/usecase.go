package usecase

import "github.com/ffajarpratama/go-chat/internal/repository"

type Usecase struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}
