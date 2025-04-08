package main

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

type fetchResultMsg struct {
	repos []repository
	err   error
}

func fetchGitHubRepos(input string) tea.Cmd {
	return func() tea.Msg {
		// Extract username from URL or use input as username
		token := input

		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)

		client := github.NewClient(tc)

		opts := &github.RepositoryListOptions{
			ListOptions: github.ListOptions{PerPage: 5},
		}

		githubRepos, _, err := client.Repositories.List(ctx, "", opts)
		if err != nil {
			return fetchResultMsg{nil, fmt.Errorf("github API returned error: %s", err)}
		}

		// Convert to our repository model
		repos := make([]repository, len(githubRepos))
		for i, repo := range githubRepos {
			r := repository{
				Name: *repo.Name,
				URL:  *repo.HTMLURL,
			}
			if repo.Description != nil {
				r.Description = *repo.Description
			}
			repos[i] = r
		}

		return fetchResultMsg{repos, nil}
	}
}
