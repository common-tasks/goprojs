package urlshortener

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var urlMap = make(map[string]string)
var mutex = &sync.Mutex{}

// start function
func Shorten(port string) {
	fmt.Println("shorten url service started ")
	http.HandleFunc("/", redirectHandler)
	http.HandleFunc("/shorten", shortURLHandler)
	fmt.Printf("shorten url service listening on port %s\n", port)
	http.ListenAndServe(port, nil)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("request hit on redirectHandler")
	shorturl := r.URL.Path[1:]
	url, ok := urlMap[shorturl]
	if !ok {
		http.Error(w, "Short url not found", http.StatusNotFound)
		return
	}
	fmt.Printf("long url is %s and short url is %s\n", url, shorturl)
	//redirect to original url
	http.Redirect(w, r, url, http.StatusSeeOther)

}

func shortURLHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	fmt.Println("request hit on shortURLHandler")
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "Missing url", http.StatusBadRequest)
		return
	}
	shortURL := ""
	surl, ok := urlMap[url]
	if !ok {
		shortURL := generateShortURL()

		urlMap[url] = shortURL
	} else {
		fmt.Printf("url already present in the map %s\n", surl)
	}

	fmt.Printf("long url is %s and short url is %s\n", url, shortURL)

	fmt.Fprintf(w, "short url for url %s is %s\n", url, shortURL)
	mutex.Unlock()
}

// initial function to generate random short url of 6 characters
func generateShortURL() string {
	mutex.Lock()
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	rand.Seed(time.Now().UnixNano())

	b := make([]rune, 6)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	mutex.Unlock()
	return string(b)

}
