package model

type PromotionSlug string

const (
	PromotionSlugMembership PromotionSlug = "membership"
	PromotionSlugPoint      PromotionSlug = "point"
)

type Promotion struct {
	ID        uint
	Slug      PromotionSlug
	Title     string
	PointRate float32
}

type Product struct {
	ID        uint
	Price     uint
	Title     string
	Promotion *Promotion
}
