package factory

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

func NewMember() Member {
	return Member{}
}

func (m Member) GetDataByID(id uint64) *Member {
	return &Member{
		ID:     id,
		Name:   "Bob",
		Status: MemberStatusActive,
		Points: 1000,
		MembershipLevel: &MembershipLevel{
			ID:         1,
			Title:      "Normal",
			PercentOff: 1,
		},
	}
}
