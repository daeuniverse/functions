package service

import (
	"daeuniverse/functions/pkg/github"
	web "daeuniverse/functions/pkg/web/client"
)

type FunctionConfiguration func(fs *FunctionService) error

type FunctionService struct {
	repo github.Repo
	web  web.Client
}

func NewFunctionService(cfgs ...FunctionConfiguration) (*FunctionService, error) {
	s := &FunctionService{}
	for _, cfg := range cfgs {
		err := cfg(s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func WithGithub(organization string, repo string) FunctionConfiguration {
	return func(s *FunctionService) error {
		s.repo = github.NewSession(organization, repo)
		return nil
	}
}

func WithWebClient() FunctionConfiguration {
	return func(s *FunctionService) error {
		s.web = web.NewClient()
		return nil
	}
}

func (s *FunctionService) FetchGeoData(filetype string) (*web.Response, error) {
	release, err := s.repo.GetLatestRelease()
	releaseDate := release.GetName()
	if err != nil {
		return nil, err
	}

	url := s.repo.FormURL(releaseDate, filetype)

	res, err := s.web.Get(url)
	if err != nil {
		return nil, err
	}

	return res, nil
}
