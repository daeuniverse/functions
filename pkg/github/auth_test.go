package github

import (
	"daeuniverse/functions/pkg/util"
	"testing"
)

func TestGithub_Auth(t *testing.T) {
	session := NewSession("techprober", "v2ray-rules-dat")
	result, err := session.GetLatestRelease()
	if err != nil {
		t.Fatal(err)
	}
	util.PrettyPrint(result)
}
