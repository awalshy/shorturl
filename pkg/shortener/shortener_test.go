package shortener

import "testing"

func TestShorten(t *testing.T) {
	s := NewShortener()
	url, err := s.Shorten("https://google.com")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if url.ID == "" {
		t.Errorf("expected id to be non-empty")
	}
}

func TestGetURL(t *testing.T) {
	s := NewShortener()
	url, err := s.Shorten("https://google.com")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	initialUrl, err := s.GetURL(url.ID, false)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if initialUrl.OriginalURL != "https://google.com" {
		t.Errorf("expected url to be https://google.com, got %s", url.OriginalURL)
	}
}