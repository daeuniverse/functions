package github

import (
	"testing"

	web "daeuniverse/functions/pkg/web/client"
)

func TestGithub_Auth(t *testing.T) {
	client := NewSession("techprober", "v2ray-rules-dat")
	result, err := client.GetLatestRelease()
	if err != nil {
		t.Fatal(err)
	}

	releaseDate := result.GetName()
	url := client.FormURL(releaseDate, "geosite.data")

	w := web.NewClient()
	res, err := w.GetAndSave(url, "geosite.dat")
	if err != nil {
		t.Fatal(err)
	}

	println(res.StatusCode())
}

func TestWorkflow_TriggerDaedPickBuild(t *testing.T) {
	client := NewSession("daeuniverse", "daed")
	result, err := client.TriggerPickBuild("main", "origin/main", "origin/main")
	if err != nil {
		t.Fatal(err)
	}
	println(result.StatusCode)
}
