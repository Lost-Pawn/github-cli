package activity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type GitHubActivity struct {
	ID 			string 	`json:"id"`
	Type 		string 	`json:"type"`
	Repo    	Repo  	`json:"repo"`
	CreatedAt 	string 	`json:"created_at"`

	Payload struct {
		Action 	string `json:"action"`
		Ref 	string `json:"ref"`
		RefType string `json:"ref_type"`

		Commits []struct {
			Message string `json:"message"`
		} `json:"commits"`

		Forkee struct {
			FullName string `json:"full_name"`
		} `json:"forkee"`

	    Issue struct {
			Title 	string `json:"title"`
			Number 	int    `json:"number"`
		} `json:"issue"`

		Comment struct {
			Body string `json:"body"`
		} `json:"comment"`

		PullRequest struct {
			Title 		string `json:"title"`
			PRNumber 	int    `json:"number"`
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

func DisplayActivities(username string, activities []GitHubActivity) error {
	if len(activities) == 0 {
		fmt.Printf("No recent activities found for user '%s'.\n", username)
		return nil
	}

	fmt.Printf("Recent GitHub activities for user '%s':\n", username)
	for _, activity := range activities {
		switch activity.Type {
			case "PushEvent":
				fmt.Printf("- Pushed to repository '%s' at %s\n", activity.Repo.Name, activity.CreatedAt)
				for _, commit := range activity.Payload.Commits {
					fmt.Printf("  Commit message: %s\n", commit.Message)
				}
			case "CreateEvent":
				fmt.Printf("- Created %s '%s' in repository '%s' at %s\n", activity.Payload.RefType, activity.Payload.Ref, activity.Repo.Name, activity.CreatedAt)
			case "DeleteEvent":
				fmt.Printf("- Deleted %s '%s' in repository '%s' at %s\n", activity.Payload.RefType, activity.Payload.Ref, activity.Repo.Name, activity.CreatedAt)
			case "WatchEvent":
				fmt.Printf("- Started watching repository '%s' at %s\n", activity.Repo.Name, activity.CreatedAt)
			case "ForkEvent":
				fmt.Printf("- Forked repository '%s' to '%s' at %s\n", activity.Repo.Name, activity.Payload.Forkee.FullName, activity.CreatedAt)
			case "IssuesEvent":
				fmt.Printf("- Issue '%s' (number %d) in repository '%s', Action: %s at %s\n", activity.Payload.Issue.Title, activity.Payload.Issue.Number, activity.Repo.Name, activity.Payload.Action, activity.CreatedAt)
			case "IssueCommentEvent":
				fmt.Printf("- Commented on issue in repository '%s' at %s\n", activity.Repo.Name, activity.CreatedAt)
				fmt.Printf("  Comment: %s\n", activity.Payload.Comment.Body)
			case "PullRequestEvent":
				fmt.Printf("- Pull request '%s' (number %d) in repository '%s', Action: %s at %s\n", activity.Payload.PullRequest.Title, activity.Payload.PullRequest.PRNumber, activity.Repo.Name, activity.Payload.Action, activity.CreatedAt)
			default:
				fmt.Printf("- %s in repository '%s' at %s\n", activity.Type, activity.Repo.Name, activity.CreatedAt)
		}
	}
	return nil
}

