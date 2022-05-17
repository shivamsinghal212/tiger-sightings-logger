package common

import (
	"github.com/BurntSushi/toml"
	"os"
)

func bindString(value *string, env string) {
	if envValue, ok := os.LookupEnv(env); ok {
		*value = envValue
	}
}

type Config struct {
	Server struct {
		Port uint64 `toml:"port"`
	} `toml:"server"`

	Database struct {
		Postgres struct {
			Host     string `toml:"pg_host"`
			Password string `toml:"pg_password"`
			User     string `toml:"pg_user"`
			Db       string `toml:"pg_db"`
		} `toml:"postgres"`
	}
}

func NewConfig(filepath string) Config {

	// Sets Environment variables from .env
	//err := godotenv.Load()
	//if err != nil {
	//	log.Println("Error loading .env file")
	//}

	//fmt.Println(strings.Join(os.Environ(), "|||"))

	var conf Config

	if _, err := toml.DecodeFile(filepath, &conf); err != nil {
		panic(err)
	}

	// Adding ENV bindings to the config.
	// If ENV value exists, then it overrides the current value.

	//bindString(&conf.Database.Postgres.Db, "PG_DB")
	//bindString(&conf.Database.Postgres.Host, "PG_HOST")
	//bindString(&conf.Database.Postgres.Password, "PG_PASSWORD")
	//bindString(&conf.Database.Postgres.User, "PG_USER")
	return conf
}
