package handler

import (
	"encoding/json"
	"github.com/conamu/job-submission-system/src/internal/pkg/logger"
	"github.com/conamu/job-submission-system/src/internal/server/pkg/job"
	"net/http"
)

func jobStatusHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	queue := job.FromContext(ctx)

	id := r.PathValue("id")
	log.Debug("called", "id", id)

	jobStatus, err := queue.GetStatus(id)
	if err != nil {
		log.With("error", err.Error()).Warn("error getting job status")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res := JobResponse{
		Id:     id,
		Status: jobStatus,
	}

	resData, err := json.Marshal(res)
	if err != nil {
		log.With("error", err.Error()).Error("error when reading response body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resData)
}
