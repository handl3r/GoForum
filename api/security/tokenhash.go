package security

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/twinj/uuid"
)

func HashToken(tokenString string) string {
	hasher := md5.New()
	hasher.Write([]byte(tokenString))
	theHash := hex.EncodeToString(hasher.Sum(nil))

	u := uuid.NewV4()
	theToken := theHash + u.String()
	return theToken
}
