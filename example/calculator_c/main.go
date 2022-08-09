package main

import (
	"bitgin/cashier"
	"bitgin/cashier/model"
	"fmt"
)

func main() {
	member := &model.Member{
		ID:                1,
		Name:              "Dio",
		Status:            model.MemberStatusActive,
		Points:            1000,
		MembershipLevelID: 1,
		MembershipLevel: &model.MembershipLevel{
			ID:    1,
			Title: "Normal",
		},
	}

	product := &model.Product{
		ID:    1,
		Price: uint(870),
		Title: "Product",
		Promotion: &model.Promotion{
			ID:    2,
			Code:  model.PromotionCodeC,
			Title: "C",
			PromotionMembershipLevel: &model.PromotionMembershipLevel{
				ID:                   1,
				PromotionID:          2,
				MembershipLevelID:    member.MembershipLevelID,
				MembershipPercentOff: 1,
				PointRate:            0.5,
			},
		},
	}

	calculator := cashier.New(product, 87)
	amount := calculator.GetAmount()
	fmt.Println(amount) // 827
}
