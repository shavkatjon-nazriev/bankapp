package card

import (
	"fmt"
	"bank/pkg/bank/types"
)

func ExamplePaymentSources() {
	cards := []types.Card {
		{
			ID: 1200,
			PAN: "0001",
			Balance: 550,
			Active: true,
		},
		{
			ID: 1201,
			PAN: "0002",
			Balance: 550,
			Active: false,
		},
		{
			ID: 1202,
			PAN: "0003",
			Balance: 0,
			Active: true,
		},
	}

	res := PaymentSources(cards)
	
	fmt.Println(res)

	//Output: [{ 0001 550}]
	
}