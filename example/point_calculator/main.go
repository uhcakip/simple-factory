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
			ID:        1,
			Slug:      model.PromotionSlugPoint,
			Title:     "C",
			PointRate: 0.5,
		},
	}

	member := &model.Member{
		ID:     1,
		Name:   "Dio",
		Status: model.MemberStatusActive,
		Points: 1000,
		MembershipLevel: &model.MembershipLevel{
			ID:         1,
			Title:      "Normal",
			PercentOff: 1,
		},
	}

	calculator := cashier.New(member, product, 87)
	amount := calculator.GetAmount()
	fmt.Println(amount) // 827
}
