package log_func

import (
	"cloud.google.com/go/logging"
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"log"
	"os"
)

// init exists purely as an entry point for GCP Cloud functions.
// https://cloud.google.com/functions/docs/writing#entry-point
func init() {
	loggingClient, err := logging.NewClient(context.Background(), os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatal("failed to create logging client", err)
	}
	logger := loggingClient.Logger("log-func", logging.RedirectAsJSON(os.Stdout))

	functions.CloudEvent("Log", LogEvent(&app{
		logger:      logger,
		logSeverity: logging.ParseSeverity(os.Getenv("LOG_SEVERITY")),
		prefix:      os.Getenv("LOG_PREFIX"),
	}))
}

type app struct {
	logger      *logging.Logger
	logSeverity logging.Severity
	prefix      string
}

func LogEvent(a *app) func(context.Context, cloudevents.Event) error {
	return func(ctx context.Context, event cloudevents.Event) error {
		return a.logger.LogSync(ctx, logging.Entry{
			Severity: a.logSeverity,
			Payload: fmt.Sprintf(
				"%s - id: %s, type: %s, source: %s, data: %s",
				a.prefix,
				event.ID(),
				event.Type(),
				event.Source(),
				string(event.Data()),
			),
		})
	}
}
