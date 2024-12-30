package shortener

import "testing"

func TestShorten(t *testing.T) {
	s := NewShortener()
	id, err := s.Shorten("https://google.com")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if id == "" {
		t.Errorf("expected id to be non-empty")
	}
}

func TestGetURL(t *testing.T) {
	s := NewShortener()
	id, err := s.Shorten("https://google.com")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	url, err := s.GetURL(id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if url != "https://google.com" {
		t.Errorf("expected url to be https://google.com, got %s", url)
	}
}