package api

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"time"
	"urlShortener/utils"
)

func shortenURLHomepage(writer http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./html/index.gohtml")
		_ = t.Execute(writer, false)
	} else {
		_ = r.ParseForm()
		result := addURL(r.Form["url"][0])
		t, _ := template.ParseFiles("./html/index.gohtml")
		_ = t.Execute(writer, result)

	}
}

func HandleRequests() {
	router := mux.NewRouter()

	router.HandleFunc("/short/{url_short}", RedirectURL).
		Methods("GET")
	router.HandleFunc("/short", handlePostReq).
		Methods("POST").
		Queries("url", "{long_url}")
	router.HandleFunc("/short/", handler)
	router.HandleFunc("/", shortenURLHomepage)
	//router.PathPrefix("/").Handler(spa)
	router.Use(loggingMiddleware)
	srv := &http.Server{
		Handler: router,
		Addr:    utils.Cfg.Server.Host + ":" + utils.Cfg.Server.Port,
		// timeouts
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func handler(w http.ResponseWriter, _ *http.Request) {
	// handles calls without an url
	http.Error(w, "Missing Arguments", http.StatusBadRequest)
}
