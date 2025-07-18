package result

import (
	"time"

	"github.com/goccy/go-json"
)

type ResultOk struct {
	Result        interface{}   `json:"result"`
	ExecutionTime string        `json:"execution_time"`
}

func NewResultOk(result interface{}, executionTime time.Duration) *ResultOk {
	return &ResultOk{
		Result:        result,
		ExecutionTime: executionTime.String(),
	}
}

func (r *ResultOk) GetJson() ([]byte, error) {
	return json.Marshal(r)
}
