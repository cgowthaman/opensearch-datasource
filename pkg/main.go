package main

import (
	"os"

	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/cgowthaman/opensearch-datasource/pkg/opensearch"
)

func main() {
	// Start listening to requests send from Grafana. This call is blocking so
	// it wont finish until Grafana shuts down the process or the plugin choose
	// to exit close down by itself
	err := datasource.Manage("grafana-opensearch-serverless-datasource", opensearch.NewOpenSearchDatasource, datasource.ManageOpts{})

	// Log any error if we could not start the plugin.
	if err != nil {
		log.DefaultLogger.Error(err.Error())
		os.Exit(1)
	}
}
