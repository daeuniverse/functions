package github

import (
	"fmt"
	"testing"

	"daeuniverse/functions/pkg/web"
)

func TestGithub_Auth(t *testing.T) {
	repo := NewSession("techprober", "v2ray-rules-dat")
	result, err := repo.GetLatestRelease()
	if err != nil {
		t.Fatal(err)
	}

	releaseDate := result.GetName()
	url := fmt.Sprintf("https://github.com/techprober/v2ray-rules-dat/releases/download/%s/geosite.dat", releaseDate)

	w := web.NewClient()
	res, err := w.GetAndSave(url, "geosite.dat")
	if err != nil {
		t.Fatal(err)
	}

	println(res.StatusCode())
}
