package entity

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type Currency string

const (
	CurrencyRUB = "RUB"
	CurrencyUSD = "USD"
	CurrencyEUR = "EUR"
)

func (c *Currency) Scan(value interface{}) error {
	asBytes, ok := value.([]byte)
	if !ok {
		return errors.New("Currency: Scan source is not []byte")
	}
	switch string(asBytes) {
	case "RUB":
		*c = CurrencyRUB
	case "USD":
		*c = CurrencyUSD
	case "EUR":
		*c = CurrencyEUR
	default:
		return fmt.Errorf("Currency: Scan: unknown value: %v", asBytes)
	}
	return nil
}

func (c Currency) Value() (driver.Value, error) {
	switch c {
	case CurrencyRUB:
		return "RUB", nil
	case CurrencyUSD:
		return "USD", nil
	case CurrencyEUR:
		return "EUR", nil
	}
	return nil, fmt.Errorf("unknown currency: %v", c)
}

// TODO: custom unmarshal for currency! (for validation )
