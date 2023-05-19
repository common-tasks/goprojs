package urlshortener

import (
	"net/http"
)

func Shorten() {
	http.HandleFunc("/", redirectHandler)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {

}
