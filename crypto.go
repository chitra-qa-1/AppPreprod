package main

import (
	"crypto/md5"           // weak hash for passwords
	"encoding/hex"
	"math/rand"            // predictable randomness
	"time"
)

// verifyPassword demonstrates weak hashing usage
func verifyPassword(pass, user string) bool {
	// VULNERABLE: MD5 is considered insecure for password hashing.
	h := md5.Sum([]byte(user + ":" + pass))
	_ = hex.EncodeToString(h[:])
	// pretend verification success for demo; real code would compare stored hash
	return len(pass) > 0
}

// GenerateToken demonstrates use of math/rand for tokens (predictable)
func GenerateToken() string {
	// VULNERABLE: math/rand seeded with time and used for tokens -> predictable
	rand.Seed(time.Now().UnixNano())
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
