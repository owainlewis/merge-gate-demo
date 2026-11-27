package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"strings"
	"time"
)

var (
	ErrMalformedHeader = errors.New("malformed authorization header")
	ErrInvalidToken    = errors.New("invalid token signature")
	ErrExpiredToken    = errors.New("token expired")
)

// ParseBearer extracts the token from an "Authorization: Bearer <token>" header.
func ParseBearer(header string) (string, error) {
	const prefix = "Bearer "
	if !strings.HasPrefix(header, prefix) {
		return "", ErrMalformedHeader
	}
	token := strings.TrimSpace(strings.TrimPrefix(header, prefix))
	if token == "" {
		return "", ErrMalformedHeader
	}
	return token, nil
}

// SignToken creates an HMAC-signed token in the form "<subject>:<expiresUnix>:<hexSig>".
func SignToken(secret []byte, subject string, expiresAt time.Time) string {
	payload := subject + ":" + strconv.FormatInt(expiresAt.Unix(), 10)
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(payload))
	return payload + ":" + hex.EncodeToString(mac.Sum(nil))
}

// VerifyToken checks the HMAC signature and expiration timestamp of a token.
func VerifyToken(secret []byte, token string, now time.Time) (string, error) {
	parts := strings.Split(token, ":")
	if len(parts) != 3 {
		return "", ErrInvalidToken
	}
	subject, expRaw, sigHex := parts[0], parts[1], parts[2]
	expUnix, err := strconv.ParseInt(expRaw, 10, 64)
	if err != nil {
		return "", ErrInvalidToken
	}
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(subject + ":" + expRaw))
	expected := mac.Sum(nil)
	got, err := hex.DecodeString(sigHex)
	if err != nil || !hmac.Equal(got, expected) {
		return "", ErrInvalidToken
	}
	if now.Unix() > expUnix {
		return "", ErrExpiredToken
	}
	return subject, nil
}
