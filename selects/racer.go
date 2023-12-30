package selects

import (
	"errors"
	"net/http"
	"time"
)

func WebsiteRacer(a, b string) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(2 * time.Second):
		return "", errors.New("took more than 2s")
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func(url string) {
		http.Get(url)
		close(ch)
	}(url)
	return ch
}
