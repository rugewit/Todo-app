package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	db *sql.DB
)

func Setup() {
	envPath := "./envs/.env"
	viper.SetConfigFile(envPath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Could not read config", err)
	}

	port := viper.Get("PORT").(string)
	user := viper.Get("USER").(string)
	password := viper.Get("PASSWORD").(string)
	dbName := viper.Get("DBNAME").(string)

	dbUrl := fmt.Sprintf("host=0.0.0.0 port=%s user=%s password=%s dbname=%s sslmode=disable",
		port, user, password, dbName)

	db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		fmt.Println("Could not connect to db", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Could not ping db", err)
	}
}
