package constants

type Status string

const (
	JOB_PENDING    Status = "PENDING"
	JOB_PROCESSING Status = "PROCESSING"
	JOB_COMPLETED  Status = "COMPLETED"
	CTX_LOGGER            = "LOGGER"
	CTX_QUEUE             = "QUEUE"
	CTX_POOL              = "POOL"
	CTX_WG                = "WG"
)
