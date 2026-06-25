package entity

import (
	"time"
)

type Article struct {
	ID         int       `json:"id"`
	Author     string    `json:"author"`
	Topic      string    `json:"topic"`
	Text       string    `json:"text"`
	PostedTime time.Time `json:"time"`
}
