package shortener

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"github.com/awalshy/shorturl/pkg/models"
)

func (s *Shortener) Shorten(longUrl string) (*models.URL, error) {
	var urlRegex = regexp.MustCompile(`^(http|https)://`)

	if !urlRegex.MatchString(longUrl) {
		return nil, fmt.Errorf("invalid URL format")
	}
	_, err := url.ParseRequestURI(longUrl)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %v", err)
	}

	now := time.Now()
	hash := sha256.Sum256([]byte(longUrl))
	id := hex.EncodeToString(hash[:])[:10]

	urlData := models.URL{
		ID:               id,
		OriginalURL:      longUrl,
		CreationDate:     now,
		ModificationDate: now,
		RedirectCount:    0,
	}
	
	err = s.Storage.SaveURL(id, map[string]string{
		"original_url": urlData.OriginalURL,
		"creation_date": urlData.CreationDate.Format(time.RFC3339),
		"modification_date": urlData.ModificationDate.Format(time.RFC3339),
		"redirect_count": strconv.Itoa(urlData.RedirectCount),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to save url: %v", err)
	}
	return &urlData, nil
}

func (s *Shortener) GetURL(id string, addCount bool) (*models.URL, error) {
	url, err := s.Storage.GetURL(id)
	if err != nil {
		return nil, err
	}

	if addCount {
		url.RedirectCount++;
	
		err = s.UpdateURL(id, url)
		if err != nil {
			return nil, fmt.Errorf("failed to update url: %v", err)
		}
	}

	return url, nil
}

func (s *Shortener) UpdateURL(id string, url *models.URL) error {
	now := time.Now()
	return shortener.Storage.SaveURL(id, map[string]string{
		"original_url": url.OriginalURL,
		"creation_date": url.CreationDate.Format(time.RFC3339),
		"modification_date": now.Format(time.RFC3339),
		"redirect_count": strconv.Itoa(url.RedirectCount),
	})
}

func (s *Shortener) DeleteURL(id string) error {
	err := s.Storage.DeleteURL(id)
	if err != nil {
		return fmt.Errorf("failed to delete url: %v", err)
	}
	return nil
}