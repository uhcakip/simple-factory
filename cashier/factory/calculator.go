package factory

type Calculator interface {
	GetAmount() uint64
}

type NormalCalculator struct {
	price uint64
}

type MembershipCalculator struct {
	price      uint64
	percentOff float32
}

type PointCalculator struct {
	price     uint64
	usePoints uint64
	pointRate float32
}

func NewCalculator(member *Member, product *Product, usePoints uint64) Calculator {
	switch product.Promotion.Slug {
	case PromotionSlugMembership:
		return &MembershipCalculator{
			price:      product.Price,
			percentOff: member.MembershipLevel.PercentOff,
		}
	case PromotionSlugPoint:
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

func (nc *NormalCalculator) GetAmount() (amount uint64) {
	return nc.price
}

func (mc *MembershipCalculator) GetAmount() (amount uint64) {
	amount = mc.price
	discount := uint64(float32(mc.price) * mc.percentOff)

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

func (pc *PointCalculator) GetAmount() (amount uint64) {
	amount = pc.price
	discount := uint64(float32(pc.usePoints) * pc.pointRate)

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
