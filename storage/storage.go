package storage

import (
	"database/sql"

	"github.com/new-task/storage/postgres"
	"github.com/new-task/storage/repo"
)

type StorageI interface {
	Task() repo.TaskService
}

type StoragePg struct {
	Db       *sql.DB
	taskrepo repo.TaskService
}

func NewStoragePg(db *sql.DB) *StoragePg {
	return &StoragePg{
		Db:       db,
		taskrepo: postgres.NewStorage(db),
	}
}

func (s StoragePg) Task() repo.TaskService {
	return s.taskrepo
}
