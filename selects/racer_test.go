package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compare speeds of urls", func(t *testing.T) {
		slowServer := createServer(20 * time.Millisecond)
		fastServer := createServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		got, _ := WebsiteRacer(slowUrl, fastUrl)
		want := fastUrl

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("return error for requests longer than 2s", func(t *testing.T) {
		slowServer := createServer(3 * time.Second)
		fastServer := createServer(4 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()
		_, err := WebsiteRacer(slowServer.URL, fastServer.URL)

		if err == nil {
			t.Error("expected error but did not receive one")
		}
	})
}

func createServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
