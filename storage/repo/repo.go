package repo

import "github.com/new-task/api/models"

type TaskService interface {
	// customer
	CreateC(c *models.Cregister) (*models.Customer, error)
	GetC(id int) (*models.Customer, error)
	UpdateC(c *models.Customer) (*models.Customer, error)
	DeleteC(id int) error
	// items
	CreateP(item *models.ItemCreate) (*models.Item, error)
	GetP(Id int) (*models.Item, error)
	UpdateP(item *models.ItemUpdate) (*models.Item, error)
	DeleteP(id int) error

	// transaction
	CreateT(tr *models.AddTransaction) (*models.Transaction, error)
	GetT(req *models.TrangetReq) (models.List, error)
	UpdateT(tr *models.TransactionUpdate) (*models.Transaction, error)
	DeleteT(id int) error
}
