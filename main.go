package main

import (
	"daeuniverse/functions/pkg/github"
	"daeuniverse/functions/pkg/util"
)

func main() {
	result := github.Auth()
	util.PrettyPrint(result)
}
