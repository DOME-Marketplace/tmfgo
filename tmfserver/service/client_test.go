package service

import (
	"net/url"
	"testing"
)

func TestURLProcessing(t *testing.T) {

	u, _ := url.Parse("https://tmf.mycredential.eu/?pepe=juan,luis")
	newURL := u.JoinPath("/path/to/resource")
	if newURL.String() != "https://tmf.mycredential.eu/path/to/resource?pepe=juan,luis" {
		t.Errorf("Expected %s, got %s", "https://tmf.mycredential.eu/path/to/resource", newURL.String())
	}

	q := newURL.Query()
	q.Set("pepe", "juan")
	newURL.RawQuery = q.Encode()
	if newURL.String() != "https://tmf.mycredential.eu/path/to/resource?pepe=juan" {
		t.Errorf("Expected %s, got %s", "https://tmf.mycredential.eu/path/to/resource", newURL.String())
	}

	ou, _ := url.Parse("https://tmf.mycredential.eu:9191/path/to/resource")
	if ou.Host != "tmf.mycredential.eu:9191" {
		t.Errorf("Expected %s, got %s", "tmf.mycredential.eu:9191", ou.Host)
	}

	uu := url.URL{
		Scheme:   "https",
		Host:     "tmf.mycredential.eu:9191",
		Path:     "/path/to/resource",
		RawQuery: "pepe=juan,luis",
	}

	if uu.String() != "https://tmf.mycredential.eu:9191/path/to/resource?pepe=juan,luis" {
		t.Errorf("Expected %s, got %s", "https://tmf.mycredential.eu:9191/path/to/resource?pepe=juan,luis", uu.String())
	}

}
