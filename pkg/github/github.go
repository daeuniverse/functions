package github

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v52/github"
	"golang.org/x/oauth2"
)

type Session interface {
	GetLatestRelease() (*github.RepositoryRelease, error)
	TriggerPickBuild(daedRef string, wingRef string, daeRef string) (*github.Response, error)
	FormURL(releaseDate string, filetype string) string
}

type session struct {
	Context      context.Context
	Client       *github.Client
	Organization string
	Repository   string
}

type Response = github.Response

func (s *session) FormURL(releaseDate string, filetype string) string {
	return fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/%s", s.Organization, s.Repository, releaseDate, filetype)
}

func NewSession(organization string, repo string) Session {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	return &session{
		Context:      ctx,
		Client:       github.NewClient(tc),
		Organization: organization,
		Repository:   repo,
	}
}
