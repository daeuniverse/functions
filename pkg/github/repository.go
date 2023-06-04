package github

import (
	"github.com/google/go-github/v52/github"
)

func (s *session) GetLatestRelease() (*github.RepositoryRelease, error) {
	result, _, err := s.Client.Repositories.GetLatestRelease(s.Context, s.Organization, s.Repository)
	if err != nil {
		return nil, err
	}
	return result, nil
}
