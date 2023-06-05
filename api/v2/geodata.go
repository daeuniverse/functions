package v2

import (
	"net/http"

	"daeuniverse/functions/service"
)

func GeodataHandler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	svc, err := service.NewFunctionService(
		service.WithGithub("techprober", "v2ray-rules-dat"),
		service.WithWebClient(),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	redirectPath, err := svc.GetLatestGeodataReleaseURL(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return 302(redirect)
	http.Redirect(w, r, redirectPath, http.StatusFound)
}
