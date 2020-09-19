package types

// Money представляет собой денежную сумму в минимальный единицах (центы, копейки, дирамы и т.д)
type Money int64

// Currency представляет код валюты
type Currency string

// Коды валют
const (
	TJS = "TJS"
	RUB = "RUB"
	USD = "USD"
)

// PAN представляет номер карты

type PAN string

type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}

type Payment struct {
	ID     int
	Amount Money
}

type PaymentSource struct {
	Type string // 'card'
	Number string // номер вида '5058 xxxx xxxx 8888'
	Balance Money
}