package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/Arthur1/mackerunn"
)

func main() {
	ctx := context.Background()

	var (
		mackerelApiKey      string
		mackerelHostID      string
		mackerelServiceName string
		runnRunbookPath     string
	)

	flag.StringVar(&mackerelHostID, "hostID", "", "Mackerel host ID to link check monitoring.")
	flag.StringVar(&mackerelServiceName, "service", "", "Mackerel service name to post service metrics.")
	flag.StringVar(&runnRunbookPath, "runbook", "", "Scenario Test Runbook of runn.")
	flag.Parse()

	mackerelApiKey = os.Getenv("MACKERUNN_MACKEREL_APIKEY")
	if mackerelHostID == "" {
		mackerelHostID = os.Getenv("MACKERUNN_MACKEREL_HOST_ID")
	}
	if mackerelServiceName == "" {
		mackerelServiceName = os.Getenv("MACKERUNN_MACKEREL_SERVICE_NAME")
	}
	if runnRunbookPath == "" {
		runnRunbookPath = os.Getenv("MACKERUNN_RUNN_RUNBOOK_PATH")
	}

	if mackerelApiKey == "" {
		log.Fatalln("MACKERUNN_MACKEREL_APIKEY is required")
	}
	if mackerelHostID == "" {
		log.Fatalln("-hostID or MACKERUNN_MACKEREL_HOST_ID is required")
	}
	if mackerelServiceName == "" {
		log.Fatalln("-service or MACKERUNN_MACKEREL_SERVICE_NAME is required")
	}
	if runnRunbookPath == "" {
		log.Fatalln("-runbook or MACKERUNN_RUNN_RUNBOOK_PATH is required")
	}

	opt := &mackerunn.RunnerOption{
		MackerelApiKey:      mackerelApiKey,
		MackerelHostID:      mackerelHostID,
		MackerelServiceName: mackerelServiceName,
		RunnRunbookPath:     runnRunbookPath,
	}
	mr := mackerunn.NewRunner(opt)
	if err := mr.Run(ctx); err != nil {
		log.Fatalln(err)
	}
}
