package hasher

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
	Hash(password string) string
}

type BCRYPTHasher struct {
	salt string
}

func NewBCRYPTHasher() *BCRYPTHasher {
	return &BCRYPTHasher{}
}

func (b *BCRYPTHasher) Hash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error generate hash for user: %v", err)
	}

	return string(bytes)
}
