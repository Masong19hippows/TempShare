package session

import (
	"fmt"

	"github.com/masong19hippows/tempShare/session/files"
)

// Sesion Defintions
type Session struct {
	// id     string
	User string
	files.Storage
	shares []share
	timer  string
}

// Function to Share session with emails and send email to users to sign up if not already
func (s *Session) Share(emails []string, group []share) {
	var temp []share
	copy(s.shares, temp)
	if len(emails) != 0 {
		for _, i := range emails {
			temp = append(temp, share{email: i})
		}
	}
	temp = append(temp, group...)
	s.shares = temp

	s.sendEmails()
}

// Send emails to all users without account
func (s *Session) sendEmails() {
	for _, i := range s.shares {
		if !(i.haveAccount()) {
			fmt.Printf("Sending email to %v\n", i.email)
		}
	}

}

// Modify Session Timeout Timer
func (s *Session) ModifyTimer(n string) error {
	s.timer = n
	return nil
}

// Creates Session
func CreateSession(user string, timer string) *Session {
	ses := &Session{User: user, timer: timer}
	ses.Storage.Init(user)
	return ses
}
