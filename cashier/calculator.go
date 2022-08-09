package cashier

import (
	"bitgin/cashier/factory"
	"bitgin/cashier/model"
)

type Calculator factory.Calculator

func New(product *model.Product, usePoints uint) Calculator {
	return factory.NewCalculator(product, usePoints)
}
