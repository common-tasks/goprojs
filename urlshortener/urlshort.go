package urlshortener

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var urlMap = make(map[string]string)
//start function
func Shorten() {
	http.HandleFunc("/", redirectHandler)
	http.HandleFunc("/shorten", shortURLHandler)
	http.ListenAndServe(":8080", nil)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shorturl := r.URL.Path[1:]
	url,ok:=urlMap[shorturl]
	if(!ok){
		http.Error(w,"Short url not found",http.StatusNotFound)
		return
	}

	//redirect to original url
	http.Redirect(w,r,url,http.StatusSeeOther)

}
func shortURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "Missing url", http.StatusBadRequest)
		return
	}
	shortURL := generateShortURL()
	urlMap[shortURL] = url
	fmt.Fprintf(w, "short url for url %s is %s\n", url, shortURL)
}
func generateShortURL() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
