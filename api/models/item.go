package models

type Item struct {
	Id        int    `json:"id"`
	ItemName  string `json:"item_name"`
	Cost      string `json:"cost"`
	Price     string `json:"price"`
	Sort      int    `json:"sort"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ItemCreate struct {
	ItemName string `json:"item_name"`
	Cost     string `json:"cost"`
	Price    string `json:"price"`
}

type ItemUpdate struct {
	Id       int    `json:"id"`
	ItemName string `json:"item_name"`
	Cost     string `json:"cost"`
	Price    string `json:"price"`
}

// transactions
type AddTransaction struct {
	CustomerId int    `json:"customer_id"`
	ItemId     int    `json:"item_id"`
	Quantity   int    `json:"quantity"`
	Price      string `json:"price"`
}

type Transaction struct {
	Id           int    `json:"id"`
	CustomerId   int    `json:"customer_id"`
	ItemId       int    `json:"item_id"`
	CustomerName string `json:"customer_name"`
	ItemName     string `json:"item_name"`
	Quantity     int    `json:"quantity"`
	Price        string `json:"price"`
	TotalPrice   string `json:"total_price"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type TransactionUpdate struct {
	Id         int    `json:"id"`
	CustomerId int    `json:"customer_id"`
	ItemId     int    `json:"item_id"`
	Quantity   int    `json:"quantity"`
	Price      string `json:"price"`
}

type TrangetReq struct {
	SearchKey string `json:"search"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
}

type Meta struct {
	TotalCount  int `json:"total_count"`
	PageCount   int `json:"page_count"`
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
}

type List struct {
	Transactions []*Transaction `json:"transaction_views"`
	Meta         Meta           `json:"_meta"`
}
