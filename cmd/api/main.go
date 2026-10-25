package main

import (
	"flag"
	"log/slog"
)

// currVer will represent the latest version. For now it is hard coded, later, this will be automatically updated.

const currVer = "1.0.0"

// config is a structure, that will contain the configuration settings for our application.
// for now it is just aq port and a current environment
type config struct {
	port int
	env  string
}

// application structure holds the dependencies for our HTTP handlers, helpers,
// and middleware. for Now, it just contains config structure copy and a logger

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	//fmt.Println("new project started")

	//config struct instance

	var cfg config

	flag.IntVar(&cfg.port, "port : ", 4000, "API Server running")
	flag.StringVar(&cfg.env, "env", "Development", "Environment(developent | staging | Production)")
	flag.Parse()
}
