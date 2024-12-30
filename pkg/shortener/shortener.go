package shortener

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"github.com/awalshy/shorturl/pkg/storage"
	"github.com/gin-gonic/gin"
)

type Shortener struct {
	Storage *storage.Storage
}

var (
	shortener *Shortener
)

func NewShortener() *Shortener {
	if shortener == nil {
		shortener = &Shortener{
			Storage: storage.GetStorage(),
		}
	}
	return shortener
}

func (s *Shortener) RegisterRoutes(r *gin.Engine) {
	r.POST("/shorten", handleSortenUrl)
	r.GET("/:id", handleUrlRedirect)
}

func handleSortenUrl(c *gin.Context) {
	var req struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortener := NewShortener()
	id, err := shortener.Shorten(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func handleUrlRedirect(c *gin.Context) {
	id := c.Param("id")

	shortener := NewShortener()
	url, err := shortener.GetURL(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}

func (s *Shortener) Shorten(longUrl string) (string, error) {
	var urlRegex = regexp.MustCompile(`^(http|https)://`)

	if !urlRegex.MatchString(longUrl) {
		return "", fmt.Errorf("invalid URL format")
	}

	_, err := url.ParseRequestURI(longUrl)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %v", err)
	}

	hash := sha256.Sum256([]byte(longUrl))
	id := hex.EncodeToString(hash[:])[:10]
	
	err = s.Storage.SaveURL(id, longUrl)
	if err != nil {
		return "", fmt.Errorf("failed to save url: %v", err)
	}
	return id, nil
}

func (s *Shortener) GetURL(id string) (string, error) {
	return s.Storage.GetURL(id)
}