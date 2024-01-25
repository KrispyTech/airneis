package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	dw "github.com/bensch777/discord-webhook-golang"
	"github.com/google/go-github/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	prs, err := getPullRequests()
	if err != nil {
		log.Fatal("Error getting pull requests")
	}

	embed := dw.Embed{
		Title:     "Open : Pull Requests",
		Color:     3450963,
		Url:       "https://github.com/KrispyTech/airneis/pulls?q=is%3Apr+is%3Aopen+review%3Arequired",
		Timestamp: time.Now(),
		Fields:    prepareMessage(prs),
	}

	if err = sendEmbed(os.Getenv("WEBHOOK_URL"), embed); err != nil {
		fmt.Println(err)
	}
}

func sendEmbed(link string, embeds dw.Embed) error {
	hook := dw.Hook{
		Username:   "Github Reminder - Airneis",
		Avatar_url: "https://github.githubassets.com/assets/GitHub-Mark-ea2971cee799.png",
		Embeds:     []dw.Embed{embeds},
	}

	payload, err := json.Marshal(hook)
	if err != nil {
		log.Fatal(err)
	}

	err = dw.ExecuteWebhook(link, payload)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func getPullRequests() ([]*github.PullRequest, error) {
	ctx := context.Background()
	client := github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)))

	prs, res, err := client.PullRequests.List(ctx, "krispyTech", "airneis", &github.PullRequestListOptions{State: "open"})
	if err != nil || res.StatusCode != http.StatusOK {
		return []*github.PullRequest{}, err
	}

	return prs, nil
}

func prepareMessage(prs []*github.PullRequest) (fields []dw.Field) {
	for _, pr := range prs {
		field := dw.Field{
			Value: fmt.Sprintf("%d: %s - %s\n", pr.GetNumber(), pr.GetTitle(), pr.GetURL()),
		}

		fields = append(fields, field)
	}

	return
}
