package models

import (
    "crypto/rand"
    "fmt"
    "math/big"
    "strings"
    "time"
)

// OTP is a value object representing a one-time password
type OTP struct {
    code      string
    createdAt time.Time
    expiresIn time.Duration
    used      bool // Added missing field
}

// GenerateOTP creates a new OTP
func GenerateOTP(now time.Time) (*OTP, error) {
    var builder strings.Builder
    for i := 0; i < 6; i++ {
        num, err := rand.Int(rand.Reader, big.NewInt(10))
        if err != nil {
            return nil, fmt.Errorf("failed to generate random number: %w", err)
        }
        builder.WriteString(fmt.Sprintf("%d", num.Int64()))
    }
    o := &OTP{
        code:      builder.String(),
        createdAt: now,
        expiresIn: 5 * time.Minute,
        used:      false,
    }
    return o, nil
}

// Valid checks if the OTP is valid
func (o OTP) Valid(target string) bool {
    if o.used || time.Since(o.createdAt) > o.expiresIn {
        return false
    }
    return o.code == target
}

// MarkAsUsed marks the OTP as used
func (o *OTP) MarkAsUsed() {
    o.used = true
}