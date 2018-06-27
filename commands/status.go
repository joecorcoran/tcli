package commands

import (
	"fmt"
	"net/url"

	"github.com/joecorcoran/tcli/req"
	"github.com/joecorcoran/tcli/utils"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli"
)

func getRepo(slug string) (int, string, error) {
	url := fmt.Sprintf("%s/repo/%s?include=repository.current_build", utils.BaseUrl, url.PathEscape(slug))
	status, body, err := req.Get(url)
	return status, body, err
}

func statusHandler(c *cli.Context) error {
	slug := c.String("repo")
	_, body, err := getRepo(slug)

	if err != nil {
		return utils.Bail("Could not fetch status.")
	}

	result := gjson.Get(body, "current_build.state")
	fmt.Fprintln(c.App.Writer, result)

	return nil
}

var Status = cli.Command{
	Name:   "status",
	Usage:  "Get the status of the latest build for this repo",
	Action: statusHandler,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "repo, r",
			Usage: "Repository slug e.g. travis-ci/travis-web",
		},
	},
}
