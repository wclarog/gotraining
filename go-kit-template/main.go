package main

import (
	"context"
	"flag"
	"fmt"
	"go-kit-template/config"
	"go-kit-template/database"
	"go-kit-template/feature"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	httpAddr := flag.String("http", ":"+config.Values.HTTP_PORT, "http listen address")

	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger,
		"service", "account",
		"time:", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller)

	_ = level.Info(logger).Log("msg", "service started")
	defer func() {
		_ = level.Info(logger).Log("msg", "service ended")
	}()

	ctx := context.Background()

	client, err := database.Connect(config.Values)
	defer func() {
		_ = client.Close()
	}()

	if err != nil {
		panic(err)
	}

	repository := feature.NewRepository(client)
	srv := feature.NewService(repository, logger)
	endpoints := feature.MakeEndpoints(srv)
	endpoints = feature.NewTXMiddleware(srv, endpoints)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("test %s", <-c)
	}()

	go func() {
		_ = logger.Log("listening on port", *httpAddr)
		var serverOptions []httptransport.ServerOption
		handler := feature.NewHandler(ctx, serverOptions, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	_ = level.Error(logger).Log("exit", <-errs)
}
