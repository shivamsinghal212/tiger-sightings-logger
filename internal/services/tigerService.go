package services

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"tigerhallProject/internal/models"
	"time"
)

func AddNewTiger(DB *gorm.DB, name string, dob string, latitude float64, longitude float64, lastSeen uint) (string, *models.Tiger){
	date, dateErr := time.Parse("2006-01-02", dob)

	if dateErr != nil {
		fmt.Println(dateErr)
		return "Incorrect Date Format", nil
	}
	lastSeenTime := time.Unix(int64(lastSeen), 0)
	obj := models.Tiger{Name: name}
	err, existingTiger := obj.CheckExistingTiger(DB)
	if err!=nil && errors.Is(err, gorm.ErrRecordNotFound){
		obj.IsActive = true
		obj.Dob = date
		obj.TigerSightings = []models.TigerSighting{{Latitude: latitude, Longitude: longitude,
			LastSeen: lastSeenTime}}
		DB.Create(&obj)
		return "New Entry Created", &obj
	}else if err!=nil {
		panic(err)
	}
	DB.First(&existingTiger)
	existingTiger.TigerSightings = []models.TigerSighting{{Latitude: latitude, Longitude: longitude,
		LastSeen: lastSeenTime}}
	DB.Save(&existingTiger)
	return fmt.Sprintf("%s already exists.", name), &existingTiger
}

func GetAllTigers(){

}

func AddNewTigerSighting(){

}

func ListAllSightings(){

}