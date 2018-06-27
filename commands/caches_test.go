package commands

import (
	"testing"

	"gopkg.in/jarcoal/httpmock.v1"
)

func TestCaches(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		"https://api.travis-ci.com/repo/joecorcoran%2Ffoo/caches",
		httpmock.NewStringResponder(200, `{"caches": [{"branch": "master"}, {"branch": "dev"}]}`),
	)

	output := run(Caches, []string{"tcli", "caches", "-r", "joecorcoran/foo"})
	if output != "master\ndev\n" {
		t.Error("Wrong output", output)
	}
}
