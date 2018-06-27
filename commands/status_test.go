package commands

import (
	"testing"

	"gopkg.in/jarcoal/httpmock.v1"
)

func TestStatus(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		"https://api.travis-ci.com/repo/joecorcoran%2Ffoo?include=repository.current_build",
		httpmock.NewStringResponder(200, `{"current_build": {"state": "passed"}}`),
	)

	output := run(Status, []string{"tcli", "status", "-r", "joecorcoran/foo"})
	if output != "passed\n" {
		t.Error("Wrong output", output)
	}
}
