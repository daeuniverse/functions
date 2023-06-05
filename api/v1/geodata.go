package v1

import (
	"bytes"
	"fmt"
	"io"
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

	res, err := svc.FetchGeoData(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.RawBody().Close()

	w.Header().Set("Content-Type", res.Header().Get("Content-Type"))
	w.Header().Set("Content-Length", res.Header().Get("Content-Length"))
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file))
	w.WriteHeader(res.StatusCode())

	buffer := make([]byte, 8)
	_, err = io.CopyBuffer(w, bytes.NewReader(res.Body()), buffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
