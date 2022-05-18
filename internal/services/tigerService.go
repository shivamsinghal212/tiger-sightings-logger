package services

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"tigerhallProject/internal/models"
	"time"
)

// Paginate reference: https://gorm.io/docs/scopes.html#pagination
func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.Query("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func logTigerSighting(DB *gorm.DB, existingTiger models.Tiger, lastSeenTime time.Time, latitude float64, longitude float64) {
	DB.First(&existingTiger)
	existingTiger.LastSeenOn = lastSeenTime
	existingTiger.TigerSightings = []models.TigerSighting{{Latitude: latitude, Longitude: longitude,
		SightDate: lastSeenTime}}
	DB.Save(&existingTiger)
}

func AddNewTiger(DB *gorm.DB, name string, dob string, latitude float64, longitude float64, lastSeen uint) (string, *models.Tiger) {
	date, dateErr := time.Parse("2006-01-02", dob)

	if dateErr != nil {
		fmt.Println(dateErr)
		return "Incorrect Date Format", nil
	}
	lastSeenTime := time.Unix(int64(lastSeen), 0)
	obj := models.Tiger{Name: name}
	err, existingTiger := obj.CheckExistingTigerByName(DB)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		obj.IsActive = true
		obj.Dob = date
		obj.LastSeenOn = lastSeenTime
		obj.TigerSightings = []models.TigerSighting{{Latitude: latitude, Longitude: longitude,
			SightDate: lastSeenTime}}
		DB.Create(&obj)
		return "New Entry Created", &obj
	} else if err != nil {
		panic(err)
	}

	logTigerSighting(DB, existingTiger, lastSeenTime, latitude, longitude)

	return fmt.Sprintf("%s already exists.", name), &existingTiger
}

func AddNewTigerSighting(DB *gorm.DB, tigerId uint, latitude float64, longitude float64, lastSeen uint) string {
	obj := models.Tiger{ID: tigerId}
	err, existingTiger := obj.CheckExistingTigerById(DB)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Sprintf("Invalid Tiger ID %d", tigerId)
	}
	lastSeenTime := time.Unix(int64(lastSeen), 0)
	logTigerSighting(DB, existingTiger, lastSeenTime, latitude, longitude)
	return fmt.Sprintf("Tiger Sighting Logged")
}

type GetAllTigersResp struct {
	TigerName  string `json:"tiger_name"`
	Dob        string `json:"dob"`
	LastSeenOn string `json:"last_seen_on"`
}

func GetAllTigers(DB *gorm.DB, c *gin.Context) []GetAllTigersResp {
	var tigers []models.Tiger
	var resp []GetAllTigersResp
	DB.Scopes(Paginate(c)).Select([]string{"name, dob, last_seen_on"}).Order("last_seen_on desc, name").Find(&tigers)
	for _, i := range tigers {
		resp = append(resp, GetAllTigersResp{
			TigerName:  i.Name,
			Dob:        i.Dob.Format("02-Jan-2006"),
			LastSeenOn: i.LastSeenOn.Format("2006-01-02 15:04:05"),
		})
	}
	if len(resp) == 0 {
		return nil
	}
	return resp
}

type GetAllTigerSightingsResp struct {
	TigerId   uint    `json:"tiger_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	SightDate string  `json:"sight_date"`
}

func GetAllSightings(DB *gorm.DB, c *gin.Context, tigerId uint) []GetAllTigerSightingsResp {
	var resp []GetAllTigerSightingsResp
	if tigerId == 0 {
		return resp
	}
	var tigerSightings []models.TigerSighting
	DB.Scopes(Paginate(c)).Where(&models.TigerSighting{TigerID: tigerId}).Order("sight_date desc").Find(&tigerSightings)
	for _, v := range tigerSightings {
		resp = append(resp, GetAllTigerSightingsResp{
			TigerId:   v.TigerID,
			Latitude:  v.Latitude,
			Longitude: v.Longitude,
			SightDate: v.SightDate.Format("2006-01-02 15:04:05"),
		})
	}
	return resp
}
