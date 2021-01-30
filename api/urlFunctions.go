package api

import "urlShortener/database"

type URL struct {
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
	//Usage    int    `json:"usage"`
}

func (u URL) save() {
	database.SetItem(u.ShortUrl, u.LongUrl)
}
