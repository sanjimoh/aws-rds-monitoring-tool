// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"aws-rds-monitoring-tool/controller"
	"aws-rds-monitoring-tool/models"
	"crypto/tls"
	"github.com/go-openapi/swag"
	"log"
	"net/http"
	"os"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"aws-rds-monitoring-tool/restapi/operations"
	"aws-rds-monitoring-tool/restapi/operations/aws_rds_monitoring_tool"
)

//go:generate swagger generate server --target ../../aws-rds-monitoring-tool --name Armt --spec ../swagger.yml

var RMC *controller.RdsMonitoringController

func configureFlags(api *operations.ArmtAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ArmtAPI) http.Handler {
	var err error

	// configure the api here
	api.ServeError = errors.ServeError
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	if RMC == nil {
		RMC, err = controller.NewRdsMonitoringController()
		if err != nil {
			log.Printf("Failed to initialize aws rds tool service, error: %s", err)
			os.Exit(1)
		}
	}

	api.AwsRdsMonitoringToolGetV1RdssHandler = aws_rds_monitoring_tool.GetV1RdssHandlerFunc(func(params aws_rds_monitoring_tool.GetV1RdssParams) middleware.Responder {
		rdss, err := RMC.MonitoringHandler.GetV1Rdss(*params.Region)
		if err != nil {
			return aws_rds_monitoring_tool.NewGetV1RdssInternalServerError().WithPayload(&models.Error{
				Code:    swag.Int64(500),
				Message: swag.String(err.Error()),
			})
		}

		return aws_rds_monitoring_tool.NewGetV1RdssOK().WithPayload(rdss)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
