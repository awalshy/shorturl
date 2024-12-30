package storage

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/awalshy/shorturl/pkg/models"
	"github.com/redis/go-redis/v9"
)

type Storage struct {
	Client *redis.Client
	Context context.Context
}

var (
	storage *Storage
)

func newStorage() *Storage {
	return &Storage{
		Context: context.Background(),
	}
}

func GetStorage() *Storage {
	if storage == nil {
		storage = newStorage()
		storage.Connect()
	}
	return storage
}

func (s *Storage) Connect() {
	s.Client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

func (s *Storage) Close() {
	s.Client.Close()
}

func (s *Storage) SaveURL(id string, value map[string]string) error {
	fmt.Printf("Saving URL: %v\n", value)
	err := s.Client.HSet(s.Context, id, value).Err()
	if err != nil {
		return fmt.Errorf("failed to save url: %v", err)
	}
	return nil
}

func (s *Storage) GetURL(id string) (*models.URL, error) {
	val, err := s.Client.HGetAll(s.Context, id).Result()
	if err != nil || len(val) == 0 {
		return nil, fmt.Errorf("failed to get url: %v", err)
	}

	creationDate, _ := time.Parse(time.RFC3339, val["creation_date"])
	modificationDate, _ := time.Parse(time.RFC3339, val["modification_date"])
	redirectCount, _ := strconv.Atoi(val["redirect_count"])

	return &models.URL{
		ID: id,
		OriginalURL: val["original_url"],
		CreationDate: creationDate,
		ModificationDate: modificationDate,
		RedirectCount: redirectCount,
	}, nil
}

func (s *Storage) DeleteURL(id string) error {
	err := s.Client.Del(s.Context, id).Err()
	if err != nil {
		return fmt.Errorf("failed to delete url: %v", err)
	}
	return nil
}