package mackerunn

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Arthur1/mackerunn/internal/scenariotest"
	"github.com/mackerelio/mackerel-client-go"
	"github.com/mattn/go-runewidth"
)

func (r *Runner) exportResultAsMetric(result *scenariotest.Result) error {
	vs := r.resultToMetricValues(result)
	return r.mackerelClient.PostServiceMetricValues(r.mackerelServiceName, vs)
}

func (r *Runner) resultToMetricValues(result *scenariotest.Result) []*mackerel.MetricValue {
	key := descToMetricKey(result.Description)
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

var metricKeyRegex = regexp.MustCompile("[a-zA-Z0-9._-]+")

func descToMetricKey(desc string) string {
	tokens := metricKeyRegex.FindAllString(desc, -1)
	if len(tokens) == 0 {
		tokens = []string{"noname"}
	}
	key := strings.Join(tokens, "_")
	return runewidth.Truncate(key, 200, "")
}
