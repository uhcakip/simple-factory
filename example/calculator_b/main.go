package main

import (
	"fmt"
	"github.com/uhcakip/simple-factory/cashier"
	"github.com/uhcakip/simple-factory/cashier/model"
)

func main() {
	member := &model.Member{
		ID:                1,
		Name:              "Dio",
		Status:            model.MemberStatusActive,
		MembershipLevelID: 2,
		MembershipLevel: &model.MembershipLevel{
			ID:    2,
			Title: "VIP 1",
		},
	}

	product := &model.Product{
		ID:    1,
		Price: uint(870),
		Title: "Product",
		Promotion: &model.Promotion{
			ID:    1,
			Code:  model.PromotionCodeB,
			Title: "B",
			PromotionMembershipLevel: &model.PromotionMembershipLevel{
				ID:                   1,
				PromotionID:          1,
				MembershipLevelID:    member.MembershipLevelID,
				MembershipPercentOff: 0.95,
			},
		},
	}

	calculator := cashier.New(product, 0)
	amount := calculator.GetAmount()
	fmt.Println(amount) // 826
}
