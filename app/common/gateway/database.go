package gateway

import (
	"clean/ent"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	*ent.Client
}

// singleton
var database = newDatabase()

func DatabaseConnect() *Database {
	return database
}

func newDatabase() *Database {
	entOptions := []ent.Option{}

	env := os.Getenv("ENV")

	if env == "dev" {
		entOptions = append(entOptions, ent.Debug())
	}

	dbConf := newDbConfig()

	datasource := fmt.Sprintf("host=%s port=%s user=%s dbname= %s password=%s sslmode=disable", dbConf.host, dbConf.port, dbConf.user, dbConf.name, dbConf.pass)

	client, err := ent.Open(dbConf.driver, datasource, entOptions...)
	if err != nil {
		log.Fatal(err)
	}

	return &Database{client}
}

type dbConfig struct {
	host   string
	port   string
	user   string
	name   string
	pass   string
	driver string
}

func newDbConfig() dbConfig {
	conf := new(dbConfig)

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	conf.host = host

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	conf.port = port

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "clean"
	}
	conf.user = user

	name := os.Getenv("DB_NAME")
	if name == "" {
		name = "clean"
	}
	conf.name = name

	pass := os.Getenv("DB_PASS")
	if pass == "" {
		pass = "clean"
	}
	conf.pass = pass

	driver := os.Getenv("DB_DRIVER")
	if driver == "" {
		driver = "postgres"
	}
	conf.driver = driver

	return *conf
}
