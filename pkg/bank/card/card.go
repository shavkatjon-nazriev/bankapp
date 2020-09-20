package card

import (
	"bank/pkg/bank/types"
)

// PaymentSources returns payments sources.
func PaymentSources(cards []types.Card) []types.PaymentSource {
	payment_source := []types.PaymentSource{}
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			payment_source = append(payment_source, types.PaymentSource{Number: string(card.PAN),Balance: card.Balance,})
	}
}
	return payment_source
}