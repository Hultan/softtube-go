package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	core "github.com/hultan/softtube/softtube.core"
)

var (
	logger   core.Logger
	config   *core.Config
	db       core.Database
	session  core.SessionIdentifier
	softtube *SoftTube
)

func main() {
	// Load the config file
	loadConfig()

	// Setup the client logging
	setupLogging()

	// Open the SoftTube database
	openDatabase()
	defer db.CloseDatabase()

	// Create a SoftTube session
	createSession()

	// Create a new application.
	softtube = new(SoftTube)
	err := softtube.StartApplication()
	if err != nil {
		logger.Log("Failed to start application!")
		logger.LogError(err)
	}

}

func loadConfig() {
	// Load config file
	config = new(core.Config)
	config.Load("main")
}

func setupLogging() {
	// Setup logging
	logger = core.NewLog(config.Client.Log)
	defer logger.Close()

	// Log start and finish
	logger.LogStart(config, "softtube client")
	defer logger.LogFinished("softtube client")
}

func openDatabase() core.Database {
	// Create the database object, and get all subscriptions
	conn := config.Connection
	db = core.New(conn.Server, conn.Port, conn.Database, conn.Username, conn.Password)
	db.OpenDatabase()
	return db
}

func createSession() {
	// Create session
	session, err := core.CreateSession(db)
	if err != nil {
		panic(err)
	}

	logger.Log(fmt.Sprintf("Session started : %s", session.Name))
}
