package services

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"tigerhallProject/internal/models"
	"time"
)

func AddNewTiger(DB *gorm.DB, name string, dob string, latitude float64, longitude float64) (string, *models.Tiger){
	date, dateErr := time.Parse("2006-01-02", dob)

	if dateErr != nil {
		fmt.Println(dateErr)
		return "Incorrect Date Format", nil
	}
	obj := models.Tiger{Name: name}
	err, existingTiger := obj.CheckExistingTiger(DB)
	if errors.Is(err, gorm.ErrRecordNotFound){
		obj.IsActive = true
		obj.Dob = date
		obj.TigerSightings = []models.TigerSighting{{Latitude: latitude, Longitude: longitude, LastSeen: time.Now()}}
		DB.Create(&obj)
		return "New Entry Created", &obj
	}
	DB.First(&existingTiger)
	existingTiger.TigerSightings = []models.TigerSighting{{Latitude: latitude, Longitude: longitude, LastSeen: time.Now()}}
	DB.Save(&existingTiger)
	return fmt.Sprintf("%s already exists.", name), &existingTiger
}

func GetAllTigers(){

}

func AddNewTigerSighting(){

}

func ListAllSightings(){

}