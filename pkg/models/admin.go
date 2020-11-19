package models

type Admin struct {
	ID       string `gorm:"default:uuid_generate_v4();primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
