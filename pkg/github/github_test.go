package github

import (
	"testing"

	web "daeuniverse/functions/pkg/web/client"
)

func TestGithub_Auth(t *testing.T) {
	repo := NewSession("techprober", "v2ray-rules-dat")
	result, err := repo.GetLatestRelease()
	if err != nil {
		t.Fatal(err)
	}

	releaseDate := result.GetName()
	url := repo.FormURL(releaseDate, "geosite.data")

	w := web.NewClient()
	res, err := w.GetAndSave(url, "geosite.dat")
	if err != nil {
		t.Fatal(err)
	}

	println(res.StatusCode())
}
