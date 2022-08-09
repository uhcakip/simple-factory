package cashier

import (
	"bitgin/cashier/factory"
)

type Calculator factory.Calculator

func New(memberID uint64, productID uint64, usePoints uint64) Calculator {
	memberFactory := factory.NewMember()
	member := memberFactory.GetDataByID(memberID)

	productFactory := factory.NewProduct()
	product := productFactory.GetDataByID(productID)

	return factory.NewCalculator(member, product, usePoints)
}
