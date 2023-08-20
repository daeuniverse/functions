package v2

import (
	"encoding/json"
	"net/http"

	"daeuniverse/functions/service"
)

type response struct {
	Result       string `json:"result"`
	WorkflowURL  string `json:"workflow_url"`
	WorkflowFile string `json:"workflow_file"`
}

func WorkflowHandler(w http.ResponseWriter, r *http.Request) {
	workflow := r.URL.Query().Get("name")
	w.Header().Set("Content-Type", "application/json")

	if workflow == "daed-pick-build" {
		svc, err := service.NewFunctionService(
			service.WithGithub("daeuniverse", "daed"),
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = svc.TriggerPickBuild(
			r.URL.Query().Get("daed"),
			r.URL.Query().Get("wing"),
			r.URL.Query().Get("dae"),
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		res := &response{Result: "ok!", WorkflowURL: "https://github.com/daeuniverse/daed/actions/workflows/pick-build.yml", WorkflowFile: "https://github.com/daeuniverse/daed/blob/main/.github/workflows/pick-build.yml"}
		_ = json.NewEncoder(w).Encode(res)
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		_, _ = w.Write([]byte(`{"error": "the given workflow cannot be trigger as it is not an available option."}`))

	}
}
