package token

import (
	"fmt"
	"github.com/o1egl/paseto"
	"time"
)

const SecretSymmetricKey = "symmetric-secret-key (size = 32)"

type Claims struct {
	UserID    int64 `json:"name"`
	ExpiresAt int64 `json:"expires_at"`
}

func GenerateToken(userID int64) (string, error) {
	payload := Claims{
		UserID:    userID,
		ExpiresAt: time.Now().Add(90 * 24 * time.Hour).Unix(),
	}

	// Create token and encrypt it
	encryptedToken, err := paseto.NewV2().Encrypt([]byte(SecretSymmetricKey), payload, nil)
	if err != nil {
		return "", fmt.Errorf("generate token failed: %s", err.Error())
	}
	return encryptedToken, nil
}
