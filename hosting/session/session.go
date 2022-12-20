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
	id string
	// User string
	*storage.Storage
	shares []share
	*sessionTimer.Timer
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
func (s *Session) ModifyTimer(t time.Duration) error {
	temp, err := sessionTimer.Create(t)
	if err != nil {
		return err
	}
	s.Timer = temp
	return nil
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
func Create(id string, timer time.Duration) *Session {
	temp := encryptAES(id)
	ses := &Session{id: temp}
	fmt.Println(ses)
	ses.Storage.Init(temp)
	return ses
}
