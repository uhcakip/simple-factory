package model

type PromotionCode string

const (
	PromotionCodeB PromotionCode = "B"
	PromotionCodeC PromotionCode = "C"
	PromotionCodeD PromotionCode = "D"
)

type Promotion struct {
	ID                       uint
	Code                     PromotionCode
	Title                    string
	PromotionMembershipLevel *PromotionMembershipLevel
}
