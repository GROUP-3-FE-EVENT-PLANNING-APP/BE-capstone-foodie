package main

import (
	"capstone/group3/config"
	"capstone/group3/factory"
	"capstone/group3/migration"
	"capstone/group3/routes"
)

func main() {
	//initiate db connection
	dbConn := config.InitDB()

	// run auto migrate from gorm
	migration.InitMigrate(dbConn)

	// initiate factory
	presenter := factory.InitFactory(dbConn)

	e := routes.New(presenter)

	e.Start(":80")

}
