package main

import (
	"log"
	"os"

	c "github.com/joecorcoran/tcli/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0.rc1"
	app.Name = "tcli"

	app.Commands = []cli.Command{
		c.Caches,
		c.Repos,
		c.Status,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
