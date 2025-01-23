package client

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/conamu/job-submission-system/src/internal/pkg/constants"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
)

type Client interface {
	CreateJob(ctx context.Context, data []byte) (string, error)
	GetJobStatus(ctx context.Context, id string) (string, error)
}
type client struct {
	jobCreateUrl *url.URL
	jobStatusUrl *url.URL
	httpClient   *http.Client
}

type ServiceResponse struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

func NewClient() Client {
	baseUrl := viper.GetString("client.baseUrl")
	jobCreatePath := viper.GetString("client.jobCreateUrl")
	jobStatusPath := viper.GetString("client.jobStatusUrl")

	jobCreateUrl, err := url.Parse(baseUrl + jobCreatePath)
	if err != nil {
		panic(err)
	}
	jobStatusUrl, err := url.Parse(baseUrl + jobStatusPath)
	if err != nil {
		panic(err)
	}

	return &client{
		jobCreateUrl: jobCreateUrl,
		jobStatusUrl: jobStatusUrl,
		httpClient:   http.DefaultClient,
	}
}

var ErrRateLimited = errors.New("rate limit reached or queue full")

func (c *client) CreateJob(ctx context.Context, inputData []byte) (string, error) {
	u := c.jobCreateUrl.String()
	body := bytes.NewReader(inputData)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, body)
	if err != nil {
		return "", errors.Wrap(err, "error while creating request")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "error while requesting job creation")
	}

	if res.StatusCode == http.StatusTooManyRequests {
		return "", ErrRateLimited
	}

	if res.StatusCode != http.StatusAccepted {
		return "", errors.New("non 202 status code for job creation request")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errors.Wrap(err, "error reading response body")
	}

	resData := &ServiceResponse{}
	err = json.Unmarshal(data, resData)
	if err != nil {
		return "", errors.Wrap(err, "error unmarshalling response body")
	}

	return resData.Id, nil
}

func (c *client) GetJobStatus(ctx context.Context, id string) (string, error) {
	u := c.jobStatusUrl.String() + id

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return "", errors.Wrap(err, "error while creating request")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "error while requesting job status")
	}

	if res.StatusCode == http.StatusTooManyRequests {
		return "", ErrRateLimited
	}

	if res.StatusCode == http.StatusTooEarly {
		return string(constants.JOB_PROCESSING), nil
	}

	if res.StatusCode != http.StatusOK {
		return "", errors.New("non 200 status code for job status request")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errors.Wrap(err, "error reading response body")
	}

	resData := &ServiceResponse{}
	err = json.Unmarshal(data, resData)
	if err != nil {
		return "", errors.Wrap(err, "error unmarshalling response body")
	}

	return resData.Status, nil
}
