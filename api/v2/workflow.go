package v2

import (
	"net/http"

	"daeuniverse/functions/service"
)

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
		_, _ = w.Write([]byte(`{"result": "ok!"}`))
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		_, _ = w.Write([]byte(`{"error": "the given workflow cannot be trigger as it is not an available option."}`))

	}
}
