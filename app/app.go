package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Create creates a new cli.App
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI Search IP"
	app.Usage = "Search IP and Server Names"
	app.Version = "0.0.1"
	

	app.Commands = []cli.Command{
		{
			Name:    "ip",
			Usage:  "Search IP",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "go.dev",
				},
			},
			Action: searchIP,
		},
	}

	return app
}



func searchIP(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}

}