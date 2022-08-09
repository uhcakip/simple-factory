package factory

import "github.com/uhcakip/simple-factory/cashier/model"

type Calculator interface {
	GetAmount() uint
}

type CalculatorA struct {
	price uint
}

type CalculatorB struct {
	price                uint
	membershipPercentOff float32
}

type CalculatorC struct {
	price     uint
	usePoints uint
	pointRate float32
}

type CalculatorD struct {
	price                uint
	usePoints            uint
	pointRate            float32
	membershipPercentOff float32
}

func NewCalculator(product *model.Product, usePoints uint) Calculator {
	switch product.Promotion.Code {
	case model.PromotionCodeB:
		return &CalculatorB{
			price:                product.Price,
			membershipPercentOff: product.Promotion.PromotionMembershipLevel.MembershipPercentOff,
		}
	case model.PromotionCodeC:
		return &CalculatorC{
			price:     product.Price,
			usePoints: usePoints,
			pointRate: product.Promotion.PromotionMembershipLevel.PointRate,
		}
	case model.PromotionCodeD:
		return &CalculatorD{
			price:                product.Price,
			usePoints:            usePoints,
			pointRate:            product.Promotion.PromotionMembershipLevel.PointRate,
			membershipPercentOff: product.Promotion.PromotionMembershipLevel.MembershipPercentOff,
		}
	default:
		return &CalculatorA{
			price: product.Price,
		}
	}
}

func (ca *CalculatorA) GetAmount() (amount uint) {
	return ca.price
}

func (cb *CalculatorB) GetAmount() (amount uint) {
	amount = uint(float32(cb.price) * cb.membershipPercentOff)
	return
}

func (cc *CalculatorC) GetAmount() (amount uint) {
	amount = cc.price
	discount := uint(float32(cc.usePoints) * cc.pointRate)

	if discount <= 0 {
		return
	}
	if discount >= amount {
		amount = 0
		return
	}

	amount -= discount
	return
}

func (cd *CalculatorD) GetAmount() (amount uint) {
	cc := &CalculatorC{
		price:     cd.price,
		usePoints: cd.usePoints,
		pointRate: cd.pointRate,
	}

	amount = cc.GetAmount()

	if amount == 0 {
		return
	}
	if cd.usePoints >= 100 {
		amount = uint(float32(amount) * cd.membershipPercentOff)
	}

	return
}
