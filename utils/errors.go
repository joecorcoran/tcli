package utils

import (
	"github.com/urfave/cli"
)

func Bail(msg string) cli.ExitCoder {
	return cli.NewExitError(msg, 1)
}
