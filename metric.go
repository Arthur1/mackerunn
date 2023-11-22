package mackerunn

import (
	"fmt"

	"github.com/Arthur1/mackerunn/internal/scenariotest"
	"github.com/mackerelio/mackerel-client-go"
)

func (r *Runner) exportResultAsMetric(result *scenariotest.Result) error {
	vs := r.resultToMetricValues(result)
	return r.mackerelClient.PostServiceMetricValues(r.mackerelServiceName, vs)
}

func (r *Runner) resultToMetricValues(result *scenariotest.Result) []*mackerel.MetricValue {
	// TODO: escape & truncate
	key := result.Description
	t := result.Timestamp.Unix()
	vs := []*mackerel.MetricValue{
		{
			Name:  fmt.Sprintf("mackerunn.execute_result.%s", key),
			Time:  t,
			Value: btoi(result.Err == nil),
		},
	}
	if result.Err != nil {
		vs = append(vs, &mackerel.MetricValue{
			Name:  fmt.Sprintf("mackerunn.elapsed_time.%s", key),
			Time:  t,
			Value: result.ElapsedTime.Seconds(),
		})
	}
	return vs
}

func btoi(b bool) int64 {
	if b {
		return 1
	}
	return 0
}
