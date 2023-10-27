package postgres

import "github.com/priyanshu360/lab-rank/dashboard/repository"

type postgres struct {
}

func NewPostgresRepo() repository.Repository {
	return &postgres{}
}

func (psql *postgres) GetUserByID(userID int) error {
	return nil
}