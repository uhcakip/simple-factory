package cashier

import (
	"bitgin/cashier/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var product = &model.Product{
	Price: uint(870),
	Promotion: &model.Promotion{
		PromotionMembershipLevel: &model.PromotionMembershipLevel{
			MembershipPercentOff: 1,
		},
	},
}

func TestNewCalculatorA(t *testing.T) {
	calculator := New(product, 0)
	assert.Equal(t, product.Price, calculator.GetAmount())
}

func TestNewCalculatorB(t *testing.T) {
	product.Promotion.Code = model.PromotionCodeB
	membershipPercentOffs := []float32{
		0.95,
		0.9,
		0.85,
	}

	expectedAmounts := []uint{826, 783, 739}
	actualAmounts := make([]uint, 3)

	for index, membershipPercentOff := range membershipPercentOffs {
		product.Promotion.PromotionMembershipLevel.MembershipPercentOff = membershipPercentOff
		calculator := New(product, 85)
		actualAmounts[index] = calculator.GetAmount()
	}

	assert.Equal(t, expectedAmounts, actualAmounts)
}

func TestNewCalculatorC(t *testing.T) {
	product.Promotion.Code = model.PromotionCodeC
	product.Promotion.PromotionMembershipLevel.PointRate = 0.8

	calculator := New(product, 83)
	assert.Equal(t, uint(804), calculator.GetAmount())
}

func TestNewCalculatorD(t *testing.T) {
	product.Promotion.Code = model.PromotionCodeD
	product.Promotion.PromotionMembershipLevel.PointRate = 0.8
	product.Promotion.PromotionMembershipLevel.MembershipPercentOff = 0.9

	calculator := New(product, 187)
	assert.Equal(t, uint(648), calculator.GetAmount()) // 648
}
