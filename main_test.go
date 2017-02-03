package bluntbe

import (
  "testing"
  "errors"
  "net/http"
  "net/http/httptest"
  "github.com/stretchr/testify/assert"
  "github.com/noobiehacker/bluntbe/internal/repo"
  //"github.com/noobiehacker/bluntbe/internal/user"
  //"fmt"
)

func TestOne(t *testing.T) {
}

type ReposTestClient struct {
	Repo []repo.Repo
	Err   error
}

func (c ReposTestClient) Get(string) ([]repo.Repo, error) {
	return c.Repo, c.Err
}

func TestGetReposHandler(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		description        string
		reposClient        *ReposTestClient
		url                string
		expectedStatusCode int
		expectedBody       string
	}{
		{
			description:        "missing argument user",
			reposClient:        &ReposTestClient{},
			url:                "/repos",
			expectedStatusCode: 400,
			expectedBody:       "MISSING_ARG_USER\n",
		}, {
			description: "error getting repos",
			reposClient: &ReposTestClient{
				Repo: []repo.Repo{},
				Err:   errors.New("fake test error"),
			},
			url:                "/user?user=fakeuser",
			expectedStatusCode: 500,
			expectedBody:       "INTERNAL_ERROR\n",
		}, {
			description: "no repos found",
			reposClient: &ReposTestClient{
				Repo: []repo.Repo{},
				Err:   nil,
			},
			url:                "/user?user=fakeuser",
			expectedStatusCode: 200,
			expectedBody:       `[]`,
		}, {
			description: "succesfull query",
			reposClient: &ReposTestClient{
				Repo: []repo.Repo{
					repo.Repo{Name: "test", Description: "a test"},
				},
				Err: nil,
			},
			url:                "/user?user=fakeuser",
			expectedStatusCode: 200,
			expectedBody:       `[{"name":"test","description":"a test"}]`,
		},
		// TODO not all cases are covered
	}

	for _, tc := range tests {
		app := &App{repo: tc.reposClient}

		req, err := http.NewRequest("GET", tc.url, nil)
		assert.NoError(err)

		w := httptest.NewRecorder()
		app.GetReposHandler(w, req)
//    fmt.Printf("%v\n", tc.expectedStatusCode)

		assert.Equal(tc.expectedStatusCode, w.Code, tc.description)
		assert.Equal(tc.expectedBody, w.Body.String(), tc.description)
	}
}
