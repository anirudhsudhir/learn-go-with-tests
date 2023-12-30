package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func TestWebsites(t *testing.T) {
	websites := []string{
		"google.com",
		"bing.com",
		"test.com",
	}
	want := map[string]bool{
		"google.com": true,
		"bing.com":   true,
		"test.com":   true,
	}

	got := WebsitesChecker(slowWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v,want %v", got, want)
	}
}

func slowWebsiteChecker(website string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkWebsites(b *testing.B) {
	sampleWebsites := make([]string, 100)
	for i := 0; i < 100; i++ {
		sampleWebsites[i] = "testing"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WebsitesChecker(slowWebsiteChecker, sampleWebsites)
	}
}
