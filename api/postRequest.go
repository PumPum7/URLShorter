package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handlePostReq(w http.ResponseWriter, r *http.Request) {
	url := mux.Vars(r)
	result := addURL(url["long_url"])
	if len(result) == 0 {
		http.Error(w, "Not a valid URL", http.StatusBadRequest)
	} else {
		urlSaved := URL{
			LongUrl:  url["long_url"],
			ShortUrl: result,
			//Usage:    0,
		}
		urlSaved.save()
		_ = json.NewEncoder(w).Encode(urlSaved)
	}
}

func addURL(url string) string {
	valid := checkURL(url)
	seq := GenName(4, 0)
	if !valid {
		return ""
	} else {
		return seq
	}
}

func checkURL(urlInput string) bool {
	if len(urlInput) == 0 {
		return false
	}
	_, err := http.Get(urlInput)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
