package httpClient

import (
	"net/http"
)

const (
	apiURL = "https://api.github.com/users/octocat/orgs"
)

//
type GithubAPIClient struct {
	Request  *http.Request
	Response *http.Response
}

func Test__GithubClient(thing string) {

}
