package activity

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type GitHubActivity struct {
	Type 	string `json:"type"`
	Repo    Repo  `json:"repo"`
	CreatedAt string `json:"created_at"`

	Payload struct {
		Action string `json:"action"`
		Ref string `json:"ref"`
		RefType string `json:"ref_type"`

		Commits []struct {
			Message string `json:"message"`
		} `json:"commits"`

		Forkee struct {
			FullName string `json:"full_name"`
		} `json:"forkee"`

	    Issue struct {
			Title string `json:"title"`
			Number int    `json:"number"`
		} `json:"issue"`

		Comment struct {
			Body string `json:"body"`
		} `json:"comment"`

		PullRequest struct {
			Title string `json:"title"`
			PRNumber int    `json:"number"`
		} `json:"pull_request"`

	} `json:"payload"`
}

type Repo struct {
	Name string `json:"name"`
}

func FetchGitHubActivities(username string) ([]GitHubActivity, error) {
	response, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", username))
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch GitHub activities: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode == 404 {
		return nil, fmt.Errorf("User '%s' not found.", username)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to fetch GitHub activities: %s", response.Status)
	}

	var activities []GitHubActivity
	err = json.NewDecoder(response.Body).Decode(&activities)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode GitHub activities: %v", err)
	}
	return activities, nil
}


