package github

import (
	"context"
	"github.com/google/go-github/v52/github"
	"golang.org/x/oauth2"
	"os"
)

type Repo interface {
	GetLatestRelease() (*github.RepositoryRelease, error)
	GetOranizationInput() string
	GetRepositoryInput() string
}

type session struct {
	Context      context.Context
	Client       *github.Client
	Organization string
	Repository   string
}

func NewSession(organization string, repo string) Repo {
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

func (s *session) GetOranizationInput() string {
	return s.Organization
}

func (s *session) GetRepositoryInput() string {
	return s.Repository
}
