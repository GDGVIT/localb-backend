package models

type Business struct {
	Base
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Type            string   `json:"type"`
	PhotoUrl        string   `json:"photo_url"`
	UserInfo        string   `json:"user_info"`
	SocialMediaLink string   `json:"social_media_link"`
	Location        Location `gorm:"embedded;embeddedPrefix:location_"`
}

type Location struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}
