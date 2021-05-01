package models

type LongUrlInput struct {
	LongUrl string `json:"long_url" binding:"required"`
}
