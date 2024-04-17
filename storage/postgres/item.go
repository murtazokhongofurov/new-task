package postgres

import (
	"github.com/new-task/api/models"
	"github.com/new-task/pkg/utils"
)

func (r storagePg) CreateP(p *models.ItemCreate) (*models.Item, error) {
	var item models.Item
	query := `
	INSERT INTO tbl_items(
		item_name, cost, price
	) VALUES($1, $2, $3) 
	RETURNING 
 	id, item_name, price, cost, created_at, updated_at`
	err := r.db.QueryRow(query, p.ItemName, utils.StringToDecimal(p.Cost), utils.StringToDecimal(p.Price)).
		Scan(&item.Id, &item.ItemName, &item.Price, &item.Cost, &CreatedAt, &UpdatedAt)
	if err != nil {
		return &item, err
	}
	item.CreatedAt = CreatedAt.Format(Layout)
	item.UpdatedAt = UpdatedAt.Format(Layout)
	return &item, nil
}

func (r storagePg) GetP(id int) (*models.Item, error) {
	var item models.Item

	query := `
	SELECT 
 		id, item_name, price, cost, created_at, updated_at
	FROM 
		tbl_items WHERE id=$1 AND deleted_at IS NULL`
	err := r.db.QueryRow(query, id).
		Scan(&item.Id, &item.ItemName, &item.Price, &item.Cost, &CreatedAt, &UpdatedAt)
	if err != nil {
		return &item, nil
	}
	item.CreatedAt = CreatedAt.Format(Layout)
	item.UpdatedAt = UpdatedAt.Format(Layout)
	return &item, nil
}

func (r storagePg) UpdateP(p *models.ItemUpdate) (*models.Item, error) {
	var item models.Item

	query := `
	UPDATE tbl_items SET
		item_name=$1, price=$2, cost=$3
	WHERE id=$4 AND deleted_at IS NULL
	RETURNING 
 		id, item_name, price, cost, created_at, updated_at`
	err := r.db.QueryRow(query, p.ItemName, utils.StringToDecimal(p.Price), utils.StringToDecimal(p.Cost), p.Id).
		Scan(&item.Id, &item.ItemName, &item.Price, &item.Cost, &CreatedAt, &UpdatedAt)
	if err != nil {
		return &item, nil
	}
	item.CreatedAt = CreatedAt.Format(Layout)
	item.UpdatedAt = UpdatedAt.Format(Layout)
	return &item, nil
}

func (r storagePg) DeleteP(id int) error {
	_, err := r.db.Exec(`UPDATE tbl_items SET deleted_at = now() WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}
