package commands

import (
	"io"
	"io/ioutil"

	"github.com/urfave/cli"
)

// Runs a command in isolation and returns the output
func run(command cli.Command, osArgs []string) string {
	r, w := io.Pipe()
	app := cli.NewApp()
	app.Writer = w
	app.Commands = []cli.Command{command}

	go func() {
		app.Run(osArgs)
		w.Close()
	}()
	stdout, _ := ioutil.ReadAll(r)
	return string(stdout)
}
