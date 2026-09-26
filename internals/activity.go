package activity

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

