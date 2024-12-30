package storage

import (
	"context"
	"fmt"

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

func (s *Storage) SaveURL(id string, value string) error {
	_, err := s.Client.Set(s.Context, id, value, 0).Result()
	if err != nil {
		return fmt.Errorf("failed to save url: %v", err)
	}
	return nil
}

func (s *Storage) GetURL(id string) (string, error) {
	val, err := s.Client.Get(s.Context, id).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get url: %v", err)
	}
	return val, nil
}