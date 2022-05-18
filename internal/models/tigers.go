package models

import (
	"gorm.io/gorm"
	"time"
)

type Tiger struct {
	ID             uint      `json:"id" gorm:"primary_key;autoIncrement"`
	Name           string    `json:"name" gorm:"index;unique" validate:"required"`
	Dob            time.Time `json:"dob" validate:"required"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      time.Time
	LastSeenOn     time.Time       `json:"last_seen_on"`
	TigerSightings []TigerSighting `json:"tiger_sightings"`
}

func (tiger Tiger) GetAllTigersSortedByLastSeen(DB *gorm.DB) {

}
func (tiger Tiger) CheckExistingTigerByName(DB *gorm.DB) (error, Tiger) {
	res := DB.Where("name = ?", tiger.Name).Take(&tiger)
	if res.Error != nil {
		return res.Error, tiger
	}
	return nil, tiger
}

func (tiger Tiger) CheckExistingTigerById(DB *gorm.DB) (error, Tiger) {
	res := DB.Where("id = ?", tiger.ID).Take(&tiger)
	if res.Error != nil {
		return res.Error, tiger
	}
	return nil, tiger
}

type TigerSighting struct {
	ID                  uint                 `json:"id" gorm:"primary_key;autoIncrement"`
	TigerID             uint                 `json:"tiger_id"`
	Latitude            float64              `json:"latitude"`
	Longitude           float64              `json:"longitude"`
	SightDate           time.Time            `json:"sight_date"`
	TigerSightingImages []TigerSightingImage `json:"tiger_sighting_images"`
}

type TigerSightingImage struct {
	ID              uint `json:"id" gorm:"primary_key;autoIncrement"`
	TigerSightingID uint `json:"tiger_sighting_id"`
}
