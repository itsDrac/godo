package tokens

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// JWTTokenizer implements the Tokenizer interface using HMAC-SHA256
type JWTTokenizer struct {
	secret        string
	tokenDuration time.Duration
}

// TokenData represents the data stored in the token
type TokenData struct {
	UserID int32
	Expiry int64
}

// NewJWTTokenizer creates a new JWT tokenizer
func NewJWTTokenizer(secret string, tokenDuration time.Duration) Tokenizer {
	if tokenDuration == 0 {
		tokenDuration = 24 * time.Hour
	}
	return &JWTTokenizer{
		secret:        secret,
		tokenDuration: tokenDuration,
	}
}

// GenerateToken creates a new token from the given data
func (j *JWTTokenizer) GenerateToken(data any) (string, error) {
	tokenData, ok := data.(TokenData)
	if !ok {
		return "", fmt.Errorf("invalid data type: expected TokenData")
	}

	// If expiry is not set, use default duration
	if tokenData.Expiry == 0 {
		tokenData.Expiry = time.Now().Add(j.tokenDuration).Unix()
	}

	// Create payload: userID:expiry
	payload := fmt.Sprintf("%d:%d", tokenData.UserID, tokenData.Expiry)

	// Create HMAC signature
	mac := hmac.New(sha256.New, []byte(j.secret))
	mac.Write([]byte(payload))
	sig := hex.EncodeToString(mac.Sum(nil))

	// Return token: payload.signature
	return fmt.Sprintf("%s.%s", payload, sig), nil
}

// ValidateToken validates and parses the token
func (j *JWTTokenizer) ValidateToken(token string) (any, error) {
	// Split token into payload and signature
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid token format")
	}

	payload, expectedSig := parts[0], parts[1]

	// Verify signature
	mac := hmac.New(sha256.New, []byte(j.secret))
	mac.Write([]byte(payload))
	actualSig := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(expectedSig), []byte(actualSig)) {
		return nil, fmt.Errorf("invalid token signature")
	}

	// Parse payload: userID:expiry
	payloadParts := strings.Split(payload, ":")
	if len(payloadParts) != 2 {
		return nil, fmt.Errorf("invalid token payload")
	}

	userID, err := strconv.ParseInt(payloadParts[0], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID in token: %w", err)
	}

	expiry, err := strconv.ParseInt(payloadParts[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid expiry in token: %w", err)
	}

	// Check if token is expired
	if time.Now().Unix() > expiry {
		return nil, fmt.Errorf("token expired")
	}

	return TokenData{
		UserID: int32(userID),
		Expiry: expiry,
	}, nil
}
