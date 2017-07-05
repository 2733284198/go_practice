package equal

type member struct {
	GroupID  int
	MemberID int
	UserID   int
	Token    string
	IP       string
}

func (m *member) IsEqual(other *member) bool {
	if m.GroupID != other.GroupID {
		return false
	}

	if m.MemberID != other.MemberID {
		return false
	}

	if m.UserID != other.UserID {
		return false
	}

	if m.Token != other.Token {
		return false
	}

	if m.IP != other.IP {
		return false
	}
	return true
}
