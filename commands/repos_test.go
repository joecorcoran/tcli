package commands

import (
	"testing"

	"gopkg.in/jarcoal/httpmock.v1"
)

func TestRepos(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		"https://api.travis-ci.com/owner/joecorcoran/repos",
		httpmock.NewStringResponder(200, `{"repositories": [{"slug": "joecorcoran/foo"}, {"slug": "joecorcoran/bar"}]}`),
	)

	output := run(Repos, []string{"tcli", "repos", "-o", "joecorcoran"})
	if output != "joecorcoran/foo\njoecorcoran/bar\n" {
		t.Error("Wrong output", output)
	}
}
