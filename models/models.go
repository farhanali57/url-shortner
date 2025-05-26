package models

import (
	"example.com/ShortUrl/db"
	"fmt"
	"github.com/google/uuid"
)

type Link struct {
	Id       string `json:"id"`
	Url      string `json:"url"`
	ShortUrl string `json:"short_url"`
}

func (u *Link) Save() error {
	query := `
INSERT INTO urls (id, url, short_url) VALUES (?, ?, ?)
`
	u.Id = uuid.New().String()

	stmt, err := db.DB.Prepare(query)
	if err != nil {

		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Id, u.Url, u.ShortUrl)
	return err
}

func GetURL(shortUrl string) (string, error) {
	var url string
	err := db.DB.QueryRow("SELECT url FROM urls WHERE short_url = ?", fmt.Sprintf("http://localhost:8080/%s", shortUrl)).Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}
