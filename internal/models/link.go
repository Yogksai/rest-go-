package models

type Link struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Url   string `json:"url" gorm:"not null"`
	Alias string `json:"alias" gorm:"not null;unique"`
}
