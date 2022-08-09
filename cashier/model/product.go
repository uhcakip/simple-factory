package model

type Product struct {
	ID          uint
	Price       uint
	Title       string
	PromotionID uint
	Promotion   *Promotion
}
