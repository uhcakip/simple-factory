package cashier

import (
	"bitgin/cashier/factory"
	"bitgin/cashier/model"
)

type Calculator factory.Calculator

func New(member *model.Member, product *model.Product, usePoints uint64) Calculator {
	return factory.NewCalculator(member, product, usePoints)
}
