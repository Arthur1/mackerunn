package mackerunn

import (
	"fmt"

	"github.com/Arthur1/mackerunn/internal/scenariotest"
	"github.com/mackerelio/mackerel-client-go"
	"github.com/mattn/go-runewidth"
)

func (r *Runner) exportResultAsCheckReport(result *scenariotest.Result, err error) error {
	report := r.resultToCheckReport(result, err)
	reports := &mackerel.CheckReports{
		Reports: []*mackerel.CheckReport{report},
	}
	return r.mackerelClient.PostCheckReports(reports)
}

func (r *Runner) resultToCheckReport(result *scenariotest.Result, err error) *mackerel.CheckReport {
	report := &mackerel.CheckReport{
		// TODO: escape & truncate
		Name:       fmt.Sprintf("mackerunn-%s", result.Description),
		OccurredAt: result.Timestamp.Unix(),
		Source:     mackerel.NewCheckSourceHost(r.mackerelHostID),
	}
	if err != nil {
		report.Status = mackerel.CheckStatusUnknown
		report.Message = truncateReportMessage(err.Error())
		return report
	}
	if result.Err != nil {
		report.Status = mackerel.CheckStatusCritical
		report.Message = truncateReportMessage(result.Err.Error())
		return report
	}
	report.Status = mackerel.CheckStatusOK
	return report
}

func truncateReportMessage(str string) string {
	return runewidth.Truncate(str, 1024, "...")
}
