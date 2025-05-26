package service

import (
	"fmt"
	"github.com/google/uuid"
)

func GenerateShortUrl(Url string) string {
	shortUrl := uuid.New().String()[:6]
	return fmt.Sprintf("http://localhost:3000/%s", shortUrl)
}
