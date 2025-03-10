package models

import "time"

type DeliveryPerson struct {
	ID        int64     `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	Latitude  float64   `gorm:"not null"`
	Longitude float64   `gorm:"not null"`
	Available bool      `gorm:"default:true"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
