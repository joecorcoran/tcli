package commands

import (
	"fmt"
	"net/url"

	"github.com/joecorcoran/tcli/req"
	"github.com/joecorcoran/tcli/utils"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli"
)

func getCachesForRepo(slug string) (int, string, error) {
	url := fmt.Sprintf("%s/repo/%s/caches", utils.BaseUrl, url.PathEscape(slug))
	status, body, err := req.Get(url)
	return status, body, err
}

func cachesHandler(c *cli.Context) error {
	owner := c.String("repo")
	_, body, err := getCachesForRepo(owner)

	if err != nil {
		return utils.Bail("Could not fetch caches.")
	}

	result := gjson.Get(body, "caches.#.branch")
	for _, branch := range result.Array() {
		fmt.Fprintln(c.App.Writer, branch.String())
	}

	return nil
}

var Caches = cli.Command{
	Name:   "caches",
	Usage:  "List all caches for this repo",
	Action: cachesHandler,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "repo, r",
			Usage: "Repository slug e.g. travis-ci/travis-web",
			Value: "travis-ci/travis-web",
		},
	},
}
