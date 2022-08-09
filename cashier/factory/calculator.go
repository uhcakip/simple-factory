package factory

import "bitgin/cashier/model"

type Calculator interface {
	GetAmount() uint
}

type NormalCalculator struct {
	price uint
}

type MembershipCalculator struct {
	price      uint
	percentOff float32
}

type PointCalculator struct {
	price     uint
	usePoints uint
	pointRate float32
}

func NewCalculator(member *model.Member, product *model.Product, usePoints uint) Calculator {
	switch product.Promotion.Slug {
	case model.PromotionSlugMembership:
		return &MembershipCalculator{
			price:      product.Price,
			percentOff: member.MembershipLevel.PercentOff,
		}
	case model.PromotionSlugPoint:
		return &PointCalculator{
			price:     product.Price,
			pointRate: product.Promotion.PointRate,
			usePoints: usePoints,
		}
	default:
		return &NormalCalculator{
			price: product.Price,
		}
	}
}

func (nc *NormalCalculator) GetAmount() (amount uint) {
	return nc.price
}

func (mc *MembershipCalculator) GetAmount() (amount uint) {
	amount = uint(float32(mc.price) * mc.percentOff)
	return
}

func (pc *PointCalculator) GetAmount() (amount uint) {
	amount = pc.price
	discount := uint(float32(pc.usePoints) * pc.pointRate)

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
