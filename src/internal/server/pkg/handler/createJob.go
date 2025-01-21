package handler

import (
	"encoding/json"
	"github.com/conamu/job-submission-system/src/internal/pkg/logger"
	"github.com/conamu/job-submission-system/src/internal/server/pkg/job"
	"io"
	"net/http"
)

type JobResponse struct {
	Id     string     `json:"id"`
	Status job.Status `json:"status,omitempty"`
}

func createJobHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	queue := job.FromContext(ctx)

	body := r.Body
	data, err := io.ReadAll(body)
	if err != nil {
		log.With("error", err.Error()).Error("error when reading response body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	j := job.CreateJob(ctx, data)
	id, err := queue.Place(j)
	if err != nil {
		log.With("error", err.Error()).Error("error placing job in queue")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := JobResponse{
		Id: id,
	}

	resData, err := json.Marshal(res)
	if err != nil {
		log.With("error", err.Error()).Error("error marshalling json response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resData)
	w.WriteHeader(http.StatusAccepted)
}
