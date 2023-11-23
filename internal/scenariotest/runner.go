package scenariotest

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/k1LoW/runn"
	"github.com/k1LoW/stopw"
)

type Runner struct {
	runnRunbookPath string
}

func NewRunner(runnRunbookPath string) *Runner {
	return &Runner{
		runnRunbookPath: runnRunbookPath,
	}
}

type Result struct {
	Err         error
	ElapsedTime time.Duration
	Description string
	Timestamp   time.Time
}

func (r *Runner) Run(ctx context.Context) (*Result, error) {
	opts := []runn.Option{
		runn.Book(r.runnRunbookPath),
	}
	o, err := runn.New(opts...)
	if err != nil {
		return nil, err
	}

	err = o.Run(ctx)
	res := o.Result()

	var (
		buf     bytes.Buffer
		profile stopw.Span
	)
	o.DumpProfile(&buf)
	json.NewDecoder(&buf).Decode(&profile)

	result := &Result{
		Err:         err,
		ElapsedTime: res.Elapsed,
		Description: res.Desc,
		Timestamp:   profile.StartedAt,
	}
	return result, nil
}
