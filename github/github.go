package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GitHubClient struct {
	*http.Client
}

func NewGitHubClient(client *http.Client) *GitHubClient {
	return &GitHubClient{Client: client}
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c GitHubClient) FetchUserByUsername(username string) (User, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := c.Get(url)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return User{}, err
	}

	var u User
	if err := json.Unmarshal(body, &u); err != nil {
		return User{}, err
	}
	return u, nil
}
