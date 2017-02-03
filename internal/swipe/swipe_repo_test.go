package swipe

import (
	"errors"
	"testing"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var c = repo.ReposClient{}

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

  //Set up the structs, basically the mock result of what the server will give back
	tests := []struct {
		description   string
		responder     httpmock.Responder
		expectedRepos []repo.Repo
		expectedError error
	}{
		{
			description:   "github api success",
			responder:     httpmock.NewStringResponder(200, `[{"name": "test", "description": "a test"}]`),
			expectedRepos: []repo.Repo{repo.Repo{Name: "test", Description: "a test"}},
			expectedError: nil,
		}, {
			description:   "github api success, no repos",
			responder:     httpmock.NewStringResponder(200, `[]`),
			expectedRepos: []repo.Repo{},
			expectedError: nil,
		}, {
			description:   "github api failure, not found",
			responder:     httpmock.NewStringResponder(404, `{"message": "not found"}`),
			expectedRepos: []repo.Repo(nil),
			expectedError: errors.New("github api: no results found"),
		},
		// TODO not all cases are covered
	}

	for _, tc := range tests {
		httpmock.RegisterResponder("GET", "https://api.github.com/users/fake/repos", tc.responder)

		r, err := c.Get("fake")

		assert.Equal(r, tc.expectedRepos, tc.description)
		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}
