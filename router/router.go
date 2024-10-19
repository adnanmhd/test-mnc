package router

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"test-mnc/handler"
	test_tahap_2 "test-mnc/test-tahap-2"
	"test-mnc/util"
)

type ListApi []handler.Api

func SetupRouter(conf util.APPConfig) *echo.Echo {

	router := setMiddleware()
	dbConn := util.SetDBConn(conf.Database)

	// registry all endpoint
	listApi := ListApi{
		test_tahap_2.NewRegisterHandler(dbConn),
		test_tahap_2.NewTopupHandler(dbConn),
		test_tahap_2.NewLoginHandler(dbConn),
		test_tahap_2.NewTopupHandler(dbConn),
		test_tahap_2.NewPaymentHandler(dbConn),
		test_tahap_2.NewTransferHandler(dbConn),
		test_tahap_2.NewTransactionsHandler(dbConn),
		test_tahap_2.NewProfileHandler(dbConn),
	}

	for _, api := range listApi {
		method, path := api.Endpoint()
		router.Add(method, path, api.Handler)
	}

	return router
}

func setMiddleware() (router *echo.Echo) {
	router = echo.New()
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${uri}, status=${status}, ${latency_human}\n",
	}))
	router.Use(middleware.BodyLimit("2M"))
	router.Use(middleware.Recover())
	router.Use(setupLogTransactionID)
	return
}

func setupLogTransactionID(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		// generate a new UUID for transaction_id
		transactionID := uuid.New().String()

		// attach the transaction_id to the request context
		c.Set("transaction_id", transactionID)

		// set up the logger with transaction_id
		logger := log.WithFields(log.Fields{
			"transaction_id": transactionID,
		})

		// attach the logger to the context for future use in handlers
		c.Set("logger", logger)

		// log the start of the request
		logger.Info("Starting request")

		return next(c)
	}
}
