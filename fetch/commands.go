package main

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

type fetchResultMsg struct {
	username string
	chunks   map[int][]repository
	err      error
}

func fetchGitHubRepos(m model) tea.Cmd {
	return func() tea.Msg {
		token := m.input

		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)

		client := github.NewClient(tc)

		// get and set username
		user, _, err := client.Users.Get(ctx, "")
		if err != nil {
			return fetchResultMsg{"", nil, fmt.Errorf("failed to get authenticated user: %s", err)}
		}
		username := *user.Login

		opts := &github.RepositoryListOptions{
			Sort:        "updated",
			Direction:   "desc",
			ListOptions: github.ListOptions{PerPage: 100},
		}

		githubRepos, _, err := client.Repositories.List(ctx, "", opts)
		if err != nil {
			return fetchResultMsg{"", nil, fmt.Errorf("github API returned error: %s", err)}
		}

		repos := make([]repository, len(githubRepos))
		chunks := make(map[int][]repository)

		for i, repo := range githubRepos {
			r := repository{
				ID:   i + 1,
				Name: *repo.Name,
				URL:  *repo.HTMLURL,
			}
			if repo.Description != nil {
				r.Description = *repo.Description
			}

			repos[i] = r
			chunkIndex := i / 10
			chunks[chunkIndex] = append(chunks[chunkIndex], r)
		}

		return fetchResultMsg{username, chunks, nil}
	}
}
