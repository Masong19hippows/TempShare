package session

// The emails the session was shared to
type share struct {
	email string
}

func (s *share) haveAccount() bool {
	return false
}
