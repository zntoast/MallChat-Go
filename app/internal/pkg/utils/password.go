package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// 这里使用简单的SHA256哈希，实际应该使用bcrypt等更安全的算法
func HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

func ValidatePassword(password, hashedPassword string) bool {
	return HashPassword(password) == hashedPassword
}
