package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GitHubEvent struct {
	Type string `json:"type"`
	Repo struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"repo"`
}

func main() {
	if len(os.Args) > 2 {
		fmt.Println("invalid command\n Usage GitHub-activity-cli <username>")
		return
	}
	username := os.Args[1]
	url := "https://api.github.com/users/" + username + "/events/public"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Print("Api not available")
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var gitEvent []GitHubEvent
	err = json.Unmarshal(data, &gitEvent)

	if len(gitEvent) == 0 {
		fmt.Println("No recent activity on ", username)
		return
	}

	for _, event := range gitEvent {

		action := ""
		switch event.Type {
		case "PushEvent":
			action = "Pushed commits to"
		case "WatchEvent":
			action = "Starred"
		case "CreateEvent":
			action = "Created a resource in"
		default:
			action = event.Type
		}

		fmt.Printf("- %s %s\n", action, event.Repo.Name)

	}

}
