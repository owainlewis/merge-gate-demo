package auth

import (
	"errors"
	"testing"
	"time"
)

func TestParseBearer(t *testing.T) {
	tok, err := ParseBearer("Bearer abc123")
	if err != nil || tok != "abc123" {
		t.Fatalf("ParseBearer = %q, %v", tok, err)
	}
	if _, err := ParseBearer("Basic xyz"); !errors.Is(err, ErrMalformedHeader) {
		t.Fatalf("expected ErrMalformedHeader, got %v", err)
	}
}

func TestSignAndVerifyToken(t *testing.T) {
	secret := []byte("demo-signing-secret")
	now := time.Unix(1_700_000_000, 0)
	token := SignToken(secret, "user_42", now.Add(time.Hour))

	subject, err := VerifyToken(secret, token, now)
	if err != nil || subject != "user_42" {
		t.Fatalf("VerifyToken = %q, %v", subject, err)
	}
	if _, err := VerifyToken(secret, token, now.Add(2*time.Hour)); !errors.Is(err, ErrExpiredToken) {
		t.Fatalf("expected ErrExpiredToken, got %v", err)
	}
	if _, err := VerifyToken([]byte("wrong"), token, now); !errors.Is(err, ErrInvalidToken) {
		t.Fatalf("expected ErrInvalidToken, got %v", err)
	}
}
