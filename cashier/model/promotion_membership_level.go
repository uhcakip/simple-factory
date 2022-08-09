package model

type PromotionMembershipLevel struct {
	ID                   uint
	PromotionID          uint
	MembershipLevelID    uint
	MembershipPercentOff float32
	PointRate            float32
}
