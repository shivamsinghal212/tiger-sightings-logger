package services

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http/httptest"

	"os"
	"testing"
	"tigerhallProject/internal/models"
)

var postgresDB *gorm.DB
var c *gin.Context

//var userService *UserService
var configFile = flag.String("config-file", "../../configs/test.toml", "Config file to use")

func TestMain(m *testing.M) {
	flag.Parse()
	var err error
	MustSetDb()
	c, _ = gin.CreateTestContext(httptest.NewRecorder())
	//logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	//userService = &UserService{
	//	DB:     postgresDB,
	//	Logger: logger,
	//}
	os.Exit(m.Run())
}

func MustSetDb() {

	var err error
	var conf struct {
		Database struct {
			Postgres struct {
				Host     string `toml:"pg_host"`
				Password string `toml:"pg_password"`
				User     string `toml:"pg_user"`
				Db       string `toml:"pg_db"`
			}
		}
	}
	if _, err := toml.DecodeFile(*configFile, &conf); err != nil {
		fmt.Printf("failed decoding file %v...\n", err)
		panic(err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		conf.Database.Postgres.Host,
		conf.Database.Postgres.User,
		conf.Database.Postgres.Password,
		conf.Database.Postgres.Db,
		5432)

	postgresDB, err = gorm.Open(postgres.New(postgres.Config{DSN: dsn}))
	if err != nil {
		fmt.Println("failed setting up postgres connection")
	}
	if err != nil {
		fmt.Printf("failed connecting %v...\n", err)
		panic(err)
	}
	err = postgresDB.AutoMigrate(&models.Tiger{})
	err = postgresDB.AutoMigrate(&models.TigerSighting{})
	err = postgresDB.AutoMigrate(&models.TigerSightingImage{})
	if err != nil {
		fmt.Printf("failed migrating %v...\n", err)
		panic(err)
	}
	//resetDB()
}

func resetDB() {
	postgresDB.Exec(`Truncate table tiger_sighting_images CASCADE;`)
	postgresDB.Exec(`Truncate table tiger_sightings CASCADE;`)
	postgresDB.Exec(`Truncate table tigers CASCADE;`)
}
