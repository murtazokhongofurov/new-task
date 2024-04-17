package postgres

import (
	"fmt"

	"github.com/new-task/api/models"
)

func (r storagePg) CreateT(req *models.AddTransaction) (*models.Transaction, error) {
	var res models.Transaction
	query := `
	INSERT INTO tbl_transaction(
		customer_id, item_id, quantity, price)
	VALUES($1, $2, $3, $4) 
	RETURNING 
		id, customer_id, item_id, quantity, price, created_at, updated_at`
	err := r.db.QueryRow(query, req.CustomerId, req.ItemId, req.Quantity, req.Price).
		Scan(&res.Id, &res.CustomerId, &res.ItemId, &res.Quantity, &res.Price, &CreatedAt, &UpdatedAt)
	if err != nil {
		return &res, err
	}
	res.CreatedAt = CreatedAt.Format(Layout)
	res.UpdatedAt = UpdatedAt.Format(Layout)
	return &res, nil
}

func (r storagePg) GetT(req *models.TrangetReq) (models.List, error) {
	var res = models.List{}
	fmt.Println(req)
	offset := (req.Page - 1) * req.Limit
	query := `
	SELECT 
		t.id,
		t.customer_id,
		t.item_id, 
		c.customer_name, 
		i.item_name, 
		i.price AS item_price, 
		t.quantity, 
		t.price AS transaction_price, 
		t.created_at, 
		t.updated_at,
		COUNT(*) OVER () AS total_count
	FROM 
		tbl_transaction t 
	JOIN 
		tbl_customer c ON t.customer_id = c.id
	JOIN 
		tbl_items i ON t.item_id = i.id 
	WHERE 
		(t.id::text ILIKE $1 OR c.customer_name ILIKE $1 OR i.item_name ILIKE $1) AND t.deleted_at IS NULL
	LIMIT $2 OFFSET $3
`
	rows, err := r.db.Query(query, "%"+req.SearchKey+"%", req.Limit, offset)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		temp := models.Transaction{}
		err = rows.Scan(&temp.Id, &temp.CustomerId, &temp.ItemId, &temp.CustomerName, &temp.ItemName, &temp.Price,
			&temp.Quantity, &temp.TotalPrice, &CreatedAt, &UpdatedAt, &res.Meta.TotalCount)
		if err != nil {
			return res, nil
		}
		temp.CreatedAt = CreatedAt.Format(Layout)
		temp.UpdatedAt = UpdatedAt.Format(Layout)
		res.Transactions = append(res.Transactions, &temp)
	}
	res.Meta.PageCount = (res.Meta.TotalCount + req.Limit - 1) / req.Limit
	res.Meta.CurrentPage = req.Page
	res.Meta.PerPage = req.Limit

	return res, nil
}

func (r storagePg) UpdateT(req *models.TransactionUpdate) (*models.Transaction, error) {
	var res models.Transaction
	query := `
	UPDATE 
		tbl_transaction SET price=$1, quantity=$2 WHERE id=$3 AND customer_id=$4 AND item_id=$5 
	RETURNING 
		id, customer_id, item_id, quantity, price, created_at, updated_at`
	err := r.db.QueryRow(query, req.Price, req.Quantity, req.Id, req.CustomerId, req.ItemId).
		Scan(&res.Id, &res.CustomerId, &res.ItemId, &res.Quantity, &res.Price, &CreatedAt, &UpdatedAt)
	if err != nil {
		return &res, err
	}
	res.CreatedAt = CreatedAt.Format(Layout)
	res.UpdatedAt = UpdatedAt.Format(Layout)
	return &res, nil
}

func (r storagePg) DeleteT(id int) error {
	_, err := r.db.Exec(`UPDATE tbl_transaction SET deleted_at=now() WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}
