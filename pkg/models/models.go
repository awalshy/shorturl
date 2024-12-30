package models

import "time"

type URL struct {
    ID               string    `json:"id"`
    OriginalURL      string    `json:"original_url"`
    CreationDate     time.Time `json:"creation_date"`
    ModificationDate time.Time `json:"modification_date"`
    RedirectCount    int       `json:"redirect_count"`
}