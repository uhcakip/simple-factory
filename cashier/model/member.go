package model

type MemberStatus uint8

const (
	MemberStatusActive   MemberStatus = 1
	MemberStatusInactive MemberStatus = 2
)

// Member type
type Member struct {
	ID              uint64
	Name            string
	Status          MemberStatus
	Points          uint
	MembershipLevel *MembershipLevel
}

// MembershipLevel type
type MembershipLevel struct {
	ID         uint64
	Title      string
	PercentOff float32
}