package utils

import (
	"log"

	"github.com/shopspring/decimal"
)

func StringToDecimal(p string) decimal.Decimal {
	price, err := decimal.NewFromString(p)
	if err != nil {
		log.Println("convert error string to decimal")
	}
	return price
}
