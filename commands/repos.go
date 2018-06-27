package commands

import (
	"fmt"

	"github.com/joecorcoran/tcli/req"
	"github.com/joecorcoran/tcli/utils"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli"
)

func getReposForOwner(owner string) (int, string, error) {
	url := fmt.Sprintf("%s/owner/%s/repos", utils.BaseUrl, owner)
	status, body, err := req.Get(url)
	return status, body, err
}

func reposHandler(c *cli.Context) error {
	owner := c.String("owner")
	_, body, err := getReposForOwner(owner)

	if err != nil {
		return utils.Bail("Could not fetch repos.")
	}

	result := gjson.Get(body, "repositories.#.slug")
	for _, slug := range result.Array() {
		fmt.Fprintln(c.App.Writer, slug.String())
	}

	return nil
}

var Repos = cli.Command{
	Name:   "repos",
	Usage:  "List all repos you can access",
	Action: reposHandler,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "owner, o",
			Usage: "Owner name",
			Value: "travis-ci",
		},
	},
}
