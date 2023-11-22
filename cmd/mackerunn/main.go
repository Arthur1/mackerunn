package main

import (
	"context"
	"os"

	"github.com/Arthur1/mackerunn"
)

func main() {
	ctx := context.TODO()
	opt := &mackerunn.RunnerOption{
		MackerelApiKey:      os.Getenv("MACKERUNN_MACKEREL_APIKEY"),
		MackerelHostID:      os.Getenv("MACKERUNN_MACKEREL_HOST_ID"),
		MackerelServiceName: os.Getenv("MACKERUNN_MACKEREL_SERVICE_NAME"),
		RunnRunBookPath:     "testdata/test.yml",
	}
	mr := mackerunn.NewRunner(opt)
	if err := mr.Run(ctx); err != nil {
		panic(err)
	}
}
