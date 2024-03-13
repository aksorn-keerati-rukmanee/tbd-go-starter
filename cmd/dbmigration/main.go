package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/configs/environment"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

func main() {

	var strFlag string
	flag.StringVar(&strFlag, "run", "none", "use \"up\" to migration up\nuse \"down\" to migration down\n")
	flag.Parse()

	environment.InitViper()
	var (
		host     = viper.Get("db.host")
		user     = viper.Get("db.user")
		password = viper.Get("db.password")
		dbname   = viper.Get("db.dbname")
		port     = viper.Get("db.port")
	)

	switch strFlag {
	case "up":
		m, err := migrate.New(
			"file://cmd/dbmigration/migrations",
			fmt.Sprintf("postgres://%v:%v@%v:%v/%v?&sslmode=disable", user, password, host, port, dbname),
		)
		if err != nil {
			log.Fatal(err)
		}

		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	case "down":
		m, err := migrate.New(
			"file://cmd/dbmigration/migrations",
			fmt.Sprintf("postgres://%v:%v@%v:%v/%v?&sslmode=disable", user, password, host, port, dbname),
		)
		if err != nil {
			log.Fatal(err)
		}

		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal(errors.New("use up or down"))
	}
}
