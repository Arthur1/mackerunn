package mackerunn

import (
	"context"

	"github.com/Arthur1/mackerunn/internal/scenariotest"
	"github.com/mackerelio/mackerel-client-go"
)

type Runner struct {
	mackerelClient      *mackerel.Client
	mackerelHostID      string
	mackerelServiceName string
	runnRunbookPath     string
}

type RunnerOption struct {
	MackerelApiKey      string
	MackerelHostID      string
	MackerelServiceName string
	RunnRunbookPath     string
}

func NewRunner(opt *RunnerOption) *Runner {
	mc := mackerel.NewClient(opt.MackerelApiKey)
	return &Runner{
		mackerelClient:      mc,
		mackerelHostID:      opt.MackerelHostID,
		mackerelServiceName: opt.MackerelServiceName,
		runnRunbookPath:     opt.RunnRunbookPath,
	}
}

func (r *Runner) Run(ctx context.Context) error {
	st := scenariotest.NewRunner(r.runnRunbookPath)
	res, _err := st.Run(ctx)
	if err := r.exportResultAsCheckReport(res, _err); err != nil {
		return err
	}
	return r.exportResultAsMetric(res)
}
