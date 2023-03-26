package app

import "github.com/urfave/cli"

// Create creates a new cli.App
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI Search IP"
	app.Usage = "Search IP and Server Names"
	app.Version = "0.0.1"
	return app
}