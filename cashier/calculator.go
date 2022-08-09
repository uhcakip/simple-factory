package cashier

import (
	"github.com/uhcakip/simple-factory/cashier/factory"
	"github.com/uhcakip/simple-factory/cashier/model"
)

type Calculator factory.Calculator

func New(product *model.Product, usePoints uint) Calculator {
	return factory.NewCalculator(product, usePoints)
}
