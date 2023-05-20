package urlshortener

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestShortenURLHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/shorten", strings.NewReader(url.Values{"url": []string{"https://facebook.com/"}}.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type","application/x-www-form-urlencoded")
	rr:=httptest.NewRecorder()
	handler:=http.HandlerFunc(shortURLHandler)
	handler.ServeHTTP(rr,req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Short URL is:"
	if rr.Body.String()[:len(expected)] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
func TestGenerateShortURL(t *testing.T) {
	url1 := generateShortURL()
	url2 := generateShortURL()

	if len(url1) != 6 {
		t.Errorf("expected length 6 but got %d\n", len(url1))
	}
	if len(url2) != 6 {
		t.Errorf("expected length 6 but got %d\n", len(url2))
	}
	if url1 == url2 {
		t.Errorf("expected unique url but got %s and %s\n", url1, url2)
	}
}
