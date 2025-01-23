package handler

import (
	"encoding/json"
	"errors"
	"github.com/conamu/job-submission-system/src/internal/pkg/constants"
	"github.com/conamu/job-submission-system/src/internal/pkg/logger"
	"github.com/conamu/job-submission-system/src/internal/pkg/schemas"
	"github.com/conamu/job-submission-system/src/internal/server/pkg/job"
	"io"
	"net/http"
)

type JobResponse struct {
	Id     string           `json:"id"`
	Status constants.Status `json:"status,omitempty"`
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

	createJobRequest := &schemas.CreateJobRequest{}
	err = json.Unmarshal(data, createJobRequest)
	if err != nil {
		log.With("error", err.Error()).Warn("error unmarshalling payload")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing or malformed request payload"))
		return
	}

	if createJobRequest.Payload == "" {
		log.Warn("empty job payload")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing request payload string"))
		return
	}

	j := job.CreateJob(ctx, createJobRequest.Payload)
	id, err := queue.Place(j)
	if err != nil {
		log.With("error", err.Error()).Error("error placing job in queue")
		if errors.Is(err, job.ErrQueueFull) {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(resData)
}
