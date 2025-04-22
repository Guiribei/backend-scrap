package utils

import (
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
)

func SanitizeCNPJ(cnpj string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(cnpj, "")
}

func SanitizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}

func IsValidEmail(email string) bool {
	return govalidator.IsEmail(email)
}