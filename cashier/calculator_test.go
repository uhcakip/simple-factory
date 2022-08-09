package cashier

import (
	"bitgin/cashier/model"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type testMembershipLevel struct{}

type testPromotion struct{}

var (
	tml testMembershipLevel
	tp  testPromotion
)

var (
	member  *model.Member
	product *model.Product
)

func TestMain(m *testing.M) {
	member = &model.Member{
		ID:     1,
		Name:   "Dio",
		Status: model.MemberStatusActive,
		Points: 1000,
	}

	product = &model.Product{
		Price: uint(870),
		Title: "Product",
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestNewNormalCalculator(t *testing.T) {
	member.MembershipLevel = tml.getLevelOne()
	product.Promotion = tp.getNormal()

	calculator := New(member, product, 100)
	assert.Equal(t, product.Price, calculator.GetAmount())
}

func TestNewMembershipCalculator(t *testing.T) {
	product.Promotion = tp.getMembership()
	membershipLevels := []*model.MembershipLevel{
		tml.getLevelOne(),
		tml.getLevelTwo(),
		tml.getLevelThree(),
	}

	expectedAmounts := []uint{826, 783, 739}
	actualAmounts := make([]uint, 3)

	for i, membershipLevel := range membershipLevels {
		member.MembershipLevel = membershipLevel
		calculator := New(member, product, 85)
		actualAmounts[i] = calculator.GetAmount()
	}

	assert.Equal(t, expectedAmounts, actualAmounts)
}

func TestNewPointCalculator(t *testing.T) {
	member.MembershipLevel = tml.getLevelOne()
	product.Promotion = tp.getPoint()
	calculator := New(member, product, 83)
	// 83 * 0.8 = 66.4 --> 66
	// 807 - 66
	assert.Equal(t, uint(804), calculator.GetAmount())
}

func (tml testMembershipLevel) getNormal() *model.MembershipLevel {
	return &model.MembershipLevel{
		ID:         1,
		Title:      "Normal",
		PercentOff: 1,
	}
}

func (tml testMembershipLevel) getLevelOne() *model.MembershipLevel {
	return &model.MembershipLevel{
		ID:         2,
		Title:      "VIP 1",
		PercentOff: 0.95,
	}
}

func (tml testMembershipLevel) getLevelTwo() *model.MembershipLevel {
	return &model.MembershipLevel{
		ID:         3,
		Title:      "VIP 2",
		PercentOff: 0.9,
	}
}

func (tml testMembershipLevel) getLevelThree() *model.MembershipLevel {
	return &model.MembershipLevel{
		ID:         3,
		Title:      "VIP 3",
		PercentOff: 0.85,
	}
}

func (tp testPromotion) getNormal() *model.Promotion {
	return &model.Promotion{
		ID:        1,
		Title:     "A",
		PointRate: 1,
	}
}

func (tp testPromotion) getMembership() *model.Promotion {
	return &model.Promotion{
		ID:    1,
		Slug:  model.PromotionSlugMembership,
		Title: "B",
	}
}

func (tp testPromotion) getPoint() *model.Promotion {
	return &model.Promotion{
		ID:        1,
		Slug:      model.PromotionSlugPoint,
		Title:     "C",
		PointRate: 0.8,
	}
}
