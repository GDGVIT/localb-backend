package models

type Business struct {
	Base
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Type            string   `json:"type"`
	PhotoUrl        string   `json:"photo_url"`
	UserInfo        string   `json:"user_info"`
	SocialMediaLink string   `json:"social_media_link"`
	Approved        bool     `json:"-"`
	Location        Location `gorm:"embedded;embeddedPrefix:location_"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GetBusinessID struct {
	BusinessID string `json:"business_id"`
}
