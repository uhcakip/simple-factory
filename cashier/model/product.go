package model

type PromotionSlug string

const (
	PromotionSlugMembership PromotionSlug = "membership"
	PromotionSlugPoint      PromotionSlug = "point"
)

type Promotion struct {
	ID        uint64
	Slug      PromotionSlug
	Title     string
	PointRate float32
}

type Product struct {
	ID        uint64
	Price     uint64
	Title     string
	Promotion *Promotion
}
