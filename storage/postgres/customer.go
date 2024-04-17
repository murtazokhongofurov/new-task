package postgres

import "github.com/new-task/api/models"
import "github.com/new-task/pkg/utils"

var res models.Customer

func (r storagePg) CreateC(cutom *models.Cregister) (*models.Customer, error) {
	query := `
	INSERT INTO tbl_customer(
		customer_name,
		balance
		) 
	VALUES(
		$1, $2
		) 
	RETURNING 
		id, customer_name, balance, created_at, updated_at`
	err := r.db.QueryRow(query, cutom.CustomerName, utils.StringToDecimal(cutom.Balance)).
		Scan(&res.Id, &res.CustomerName, &res.Balance, &CreatedAt, &UpdatedAt)
	if err != nil {
		return &res, err
	}
	res.CreatedAt = CreatedAt.Format(Layout)
	res.UpdatedAt = UpdatedAt.Format(Layout)
	return &res, nil
}

func (r storagePg) GetC(id int) (*models.Customer, error) {
	query := `
	SELECT
		id, customer_name, 
		balance, created_at, updated_at
	FROM tbl_customer WHERE id=$1 AND deleted_at IS NULL`
	err := r.db.QueryRow(query, id).
		Scan(&res.Id, &res.CustomerName, &res.Balance, &CreatedAt, &UpdatedAt)
	if err != nil {
		return &res, err
	}
	res.CreatedAt = CreatedAt.Format(Layout)
	res.UpdatedAt = UpdatedAt.Format(Layout)
	return &res, nil
}

func (r storagePg) UpdateC(cutom *models.Customer) (*models.Customer, error) {
	query := `
	UPDATE tbl_customer SET
		customer_name=$1, 
		balance=$2
	WHERE 
		id=$3 AND deleted_at IS NULL 
	RETURNING 
		id, customer_name, balance, created_at, updated_at`
	err := r.db.QueryRow(query, cutom.CustomerName, cutom.Balance, cutom.Id).
		Scan(&res.Id, &res.CustomerName, &res.Balance, &CreatedAt, &UpdatedAt)
	if err != nil {
		return &res, err
	}
	res.CreatedAt = CreatedAt.Format(Layout)
	res.UpdatedAt = UpdatedAt.Format(Layout)
	return &res, nil
}

func (r storagePg) DeleteC(id int) error {
	_, err := r.db.Exec(`UPDATE tbl_customer SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL`, id)
	if err != nil {
		return err
	}
	return nil
}
