package auth

import (
	"errors"
	"testing"
	"time"
)

func TestParseBearer(t *testing.T) {
	tests := []struct {
		name    string
		header  string
		want    string
		wantErr error
	}{
		{
			name:   "standard casing",
			header: "Bearer abc123",
			want:   "abc123",
		},
		{
			name:   "lowercase prefix",
			header: "bearer abc123",
			want:   "abc123",
		},
		{
			name:   "uppercase prefix",
			header: "BEARER abc123",
			want:   "abc123",
		},
		{
			name:   "mixed case prefix preserves token case",
			header: "BeArEr AbC123",
			want:   "AbC123",
		},
		{
			name:   "trims surrounding whitespace around token",
			header: "bearer   abc123  ",
			want:   "abc123",
		},
		{
			name:    "different scheme",
			header:  "Basic xyz",
			wantErr: ErrMalformedHeader,
		},
		{
			name:    "empty header",
			header:  "",
			wantErr: ErrMalformedHeader,
		},
		{
			name:    "prefix without space",
			header:  "Bearer",
			wantErr: ErrMalformedHeader,
		},
		{
			name:    "prefix without separator",
			header:  "Bearerabc123",
			wantErr: ErrMalformedHeader,
		},
		{
			name:    "empty token after prefix",
			header:  "bearer   ",
			wantErr: ErrMalformedHeader,
		},
		{
			name:    "leading whitespace before scheme",
			header:  " Bearer abc123",
			wantErr: ErrMalformedHeader,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tok, err := ParseBearer(tc.header)
			if tc.wantErr != nil {
				if !errors.Is(err, tc.wantErr) {
					t.Fatalf("ParseBearer(%q) error = %v, want %v", tc.header, err, tc.wantErr)
				}
				return
			}
			if err != nil || tok != tc.want {
				t.Fatalf("ParseBearer(%q) = %q, %v; want %q, nil", tc.header, tok, err, tc.want)
			}
		})
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
