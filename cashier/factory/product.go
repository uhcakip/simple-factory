package factory

type PromotionSlug string

const (
	PromotionSlugMembership PromotionSlug = "membership"
	PromotionSlugPoint      PromotionSlug = "point"
)

type Promotion struct {
	ID        uint64
	Slug      PromotionSlug
	Title     string
	PointRate float32 // e.g., 1:1
}

type Product struct {
	ID        uint64
	Price     uint64
	Title     string
	Promotion *Promotion
}

func NewProduct() Product {
	return Product{}
}

func (p Product) GetDataByID(id uint64) *Product {
	return &Product{
		ID:    id,
		Price: 870,
		Title: "商品名稱",
		Promotion: &Promotion{
			ID:        2,
			Slug:      PromotionSlugPoint,
			Title:     "Point Promotion",
			PointRate: 1,
		},
	}
}
