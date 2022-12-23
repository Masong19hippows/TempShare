package session

import (
	"crypto/aes"
	"crypto/cipher"
	cryptoRand "crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/masong19hippows/TempShare/hosting/session/sessionTimer"
	"github.com/masong19hippows/TempShare/hosting/session/storage"
)

// Sesion Defintions
type Session struct {
	ID string
	// User string
	Storage             *storage.Storage
	Shares              []share
	*sessionTimer.Timer `json:"-"`
}

// Function to Share session with emails and send email to users to sign up if not already
func (s *Session) Share(emails []string, group []share) {
	var temp []share
	copy(s.Shares, temp)
	if len(emails) != 0 {
		for _, i := range emails {
			temp = append(temp, share{email: i})
		}
	}
	temp = append(temp, group...)
	s.Shares = temp

	s.sendEmails()
}

// Send emails to all users without account
func (s *Session) sendEmails() {
	for _, i := range s.Shares {
		if !(i.haveAccount()) {
			fmt.Printf("Sending email to %v\n", i.email)
		}
	}

}

// Modify Session Timeout Timer
func (s *Session) ModifyTimer(t time.Duration) error {
	err := s.Timer.Modify(t)
	return err
}

// Encrypt given string in AES with random key. We don't want the user account stored on the Server, thats dumb
func encryptAES(plaintext string) string {
	// create cipher
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var key string
	for i := 0; i < 32; i++ {
		key += string(charset[rand.Intn(len(charset))])
	}
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(cryptoRand.Reader, nonce); err != nil {
		panic(err)
	}

	return hex.EncodeToString(gcm.Seal(nonce, nonce, []byte(plaintext), nil))
}

// Creates Session
func Create(id string, timer time.Duration) (*Session, error) {
	temp := encryptAES(id)
	ses := &Session{ID: temp, Storage: &storage.Storage{}, Timer: &sessionTimer.Timer{}}
	ses.Storage.Init(temp)
	err := ses.Timer.Init(timer)
	if err != nil {
		return nil, err
	}
	// defer ses.Delete()
	go func() {
		<-ses.Timer.Execute
		ses.Delete()
	}()
	return ses, nil
}

// Deleting and nilling Session
func (s *Session) Delete() {
	err := s.Storage.Delete()
	if err != nil {
		panic(err)
	}
	s.ID = ""
	s.Storage = nil
	s.Shares = nil
	s.Timer = nil
}
