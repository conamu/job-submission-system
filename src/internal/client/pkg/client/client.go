package client

type Client interface {
	CreateJob(data string) (string, error)
	GetJobStatus(id string) (string, error)
}
type client struct {
	baseUrl      string
	jobCreateUrl string
	jobStatusUrl string
}

func NewClient() Client {
	return &client{}
}

func (c *client) CreateJob(data string) (string, error) {

}

func (c *client) GetJobStatus(id string) (string, error) {

}
