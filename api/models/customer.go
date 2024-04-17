package models

type Cregister struct {
	CustomerName string `json:"customer_name"`
	Balance      string `json:"balance"`
}

type Customer struct {
	Id           int    `json:"id"`
	CustomerName string `json:"customer_name"`
	Balance      string `json:"balance"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
