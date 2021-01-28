package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"urlShortener/database"
	"urlShortener/utils"
)

type URL struct {
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
	Usage    int    `json:"usage"`
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	url := mux.Vars(r)
	redirectURL := database.GetItem(url["url_short"])
	if len(redirectURL) < 1 {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
	http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
	utils.Log.Infof("redirect url to %s", redirectURL)
}
