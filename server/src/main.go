package main

import (
	"comparisonLaboratories/src/core"
	"comparisonLaboratories/src/db"
	"comparisonLaboratories/src/herr"
	"comparisonLaboratories/src/transport"
	"github.com/go-pg/pg/v10"
	"os"
)

func main() {
	core.InitEnv()
	core.InitServer(core.Server)
	core.InitConfig()
	transport.InitRouters(core.Server)

	err := db.ConnectToDB(&pg.Options{
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Database: os.Getenv("DATABASE"),
	})

	herr.HandlerError(err, "Fail connect to database")
	defer func(Database *pg.DB) {
		err := Database.Close()
		herr.HandlerError(err, "Unable to close database connection")

	}(db.Database)

	port := os.Getenv("PORT")
	herr.HandlerError(core.Server.Run(port), "Server did not start")
}
