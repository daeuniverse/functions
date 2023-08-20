package github

import (
	"github.com/google/go-github/v52/github"
)

const (
	pickBuildWorkflowFile = "pick-build.yml"
)

func (s *session) TriggerPickBuild(daedRef string, wingRef string, daeRef string) (*github.Response, error) {
	triggerReq := &github.CreateWorkflowDispatchEventRequest{
		Ref:    "main",
		Inputs: map[string]interface{}{"daed": daedRef, "wing": wingRef, "dae-core": daeRef},
	}
	result, err := s.Client.Actions.CreateWorkflowDispatchEventByFileName(s.Context, s.Organization, s.Repository, pickBuildWorkflowFile, *triggerReq)
	if err != nil {
		return nil, err
	}
	return result, nil
}
