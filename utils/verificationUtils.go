package utils

import (
    "math/rand"
    "time"
)

func GenerateVerificationCode(length int) string {
    rand.Seed(time.Now().UnixNano())
    charset := "0123456789"
    result := make([]byte, length)
    for i := range result {
        result[i] = charset[rand.Intn(len(charset))]
    }
    return string(result)
}