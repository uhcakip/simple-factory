package model

type MemberStatus uint8

const (
	MemberStatusActive   MemberStatus = 1
	MemberStatusInactive MemberStatus = 2
)

// Member type
type Member struct {
	ID                uint
	Name              string
	Status            MemberStatus
	Points            uint
	MembershipLevelID uint
	MembershipLevel   *MembershipLevel
}
