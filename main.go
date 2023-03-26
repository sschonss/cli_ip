package main

import (
	"cli_ip/app"
	"fmt"
	"log"
	"os"
)

func main() { 
	fmt.Println("Welcome to the CLI Search IP")
	
	// Create a new app
	application := app.Generate()
	err := application.Run(os.Args); if err != nil {
		log.Fatal(err)
	}

	

}
