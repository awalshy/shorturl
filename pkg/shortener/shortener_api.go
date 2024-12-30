package shortener

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (s *Shortener) RegisterRoutes(r *gin.Engine) {
	r.POST("/shorten", handleSortenUrl)
	r.PATCH("/:id", handleUpdateUrl)
	r.GET("/:id", handleUrlRedirect)
	r.GET("/:id/stats", handleGetUrlStats)
	r.DELETE("/:id", handleDeleteUrl)
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
	url, err := shortener.Shorten(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": url.ID,
		"longUrl": url.OriginalURL,
		"shortUrl": fmt.Sprintf("%s/%s", c.Request.Host, url.ID),
		"creationDate": url.CreationDate,
		"modificationDate": url.ModificationDate,
	})
}

func handleUrlRedirect(c *gin.Context) {
	id := c.Param("id")

	shortener := NewShortener()
	url, err := shortener.GetURL(id,  true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}

func handleUpdateUrl(c *gin.Context) {
	id := c.Param("id")

	var req struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortener := NewShortener()
	url, err := shortener.GetURL(id, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	url.OriginalURL = req.URL
	err = shortener.UpdateURL(id, url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
		"longUrl": req.URL,
		"shortUrl": fmt.Sprintf("%s/%s", c.Request.Host, id),
	})
}

func handleDeleteUrl(c *gin.Context) {
	id := c.Param("id")

	shortener := NewShortener()
	err := shortener.DeleteURL(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "URL deleted"})
}

func handleGetUrlStats(c *gin.Context) {
	id := c.Param("id")

	shortener := NewShortener()
	url, err := shortener.GetURL(id, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
		"longUrl": url.OriginalURL,
		"shortCode": id,
		"creationDate": url.CreationDate,
		"modificationDate": url.ModificationDate,
		"redirectCount": url.RedirectCount,
	})
}