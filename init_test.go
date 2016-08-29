package gitlab

import (
	"log"
	"os"
	"strconv"
)

var (
	PersonalToken string
	ServerURL     string
	APIPath       string
	RepoPath      string
	RepoID        int
)

func init() {
	if PersonalToken = os.Getenv("PERSONAL_TOKEN"); PersonalToken == "" {
		log.Fatal("You must set environment variable PERSONAL_TOKEN to run tests")
	}

	if ServerURL = os.Getenv("SERVER_URL"); ServerURL == "" {
		log.Fatal("You must set environment variable SERVER_URL to run tests")
	}

	if APIPath = os.Getenv("API_PATH"); APIPath == "" {
		APIPath = "/api/v3"
	}

	if RepoPath = os.Getenv("REPO_PATH"); RepoPath == "" {
		log.Fatal("You must set environment variable REPO_PATH to run tests, the reposotory must a copy of https://gitlab.com/Ronmi/test-project")
	}

	var err error
	RepoID, err = strconv.Atoi(os.Getenv("REPO_ID"))
	if err != nil {
		log.Fatal(`You must set environment variable REPO_ID to run tests, repository id can be found at "Triggers" page`)
	}
}

// this always return PAT client, since oauth client is not possible due to auth flow
func makeClient() *GitLab {
	return FromPAT(ServerURL, APIPath, PersonalToken, nil)
}
