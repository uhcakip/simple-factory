package main

import (
	"bitgin/cashier"
	"bitgin/cashier/model"
	"fmt"
)

func main() {
	product := &model.Product{
		ID:    1,
		Price: uint64(870),
		Title: "Product",
		Promotion: &model.Promotion{
			ID:    1,
			Slug:  model.PromotionSlugMembership,
			Title: "B",
		},
	}

	member := &model.Member{
		ID:     1,
		Name:   "Dio",
		Status: model.MemberStatusActive,
		MembershipLevel: &model.MembershipLevel{
			ID:         2,
			Title:      "VIP 1",
			PercentOff: 0.95,
		},
	}

	calculator := cashier.New(member, product, 0)
	amount := calculator.GetAmount()
	fmt.Println(amount) // 826
}
