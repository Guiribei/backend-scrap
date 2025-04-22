package middleware

import (
	"net/http"
	"os"
)

func APIKeyMiddleware(next http.Handler) http.Handler {
	expectedKey := os.Getenv("API_KEY")
	if expectedKey == "" {
		panic("API_KEY variable is not set in this environment")
	}
}