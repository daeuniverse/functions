package main

import (
	"log"

	"daeuniverse/functions/service"
)

func main() {
	svc, err := service.NewFunctionService(
		service.WithGithub("techprober", "v2ray-rules-dat"),
		service.WithWebClient(),
	)
	if err != nil {
		log.Fatal(err)
	}

	result, err := svc.FetchGeoData("geosite.dat")
	if err != nil {
		log.Fatal(err)
	}

	println(result.StatusCode())
}
