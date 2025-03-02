package models

import "time"

type Order struct {
	ID               int64     `gorm:"primaryKey"`
	DeliveryPersonID int64     `gorm:"not null"`
	Status           string    `gorm:"not null"`
	AssignedAt       time.Time `gorm:"autoCreateTime"`
}
