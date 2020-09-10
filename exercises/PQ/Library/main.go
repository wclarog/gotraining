package main

import (
	"context"
	"excercise-library/config"
	"excercise-library/database"
	"excercise-library/materials"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	/*
		connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			config.DB.DB_USER,
			config.DB.DB_PASS,
			config.DB.DB_HOST,
			config.DB.DB_PORT,
			config.DB.DB_NAME)

		client, errOpen := ent.Open("mysql", connectionString)
	*/

	client, err := database.Connect(config.Values)
	if err != nil {
		panic("database connection failed")
	}
	defer client.Close()

	repository := materials.NewRepository(client)
	srv := materials.NewService(repository, logger)
	endpoints := materials.MakeEndpoints(srv)
	endpoints = materials.NewAuthMiddleware(endpoints)
	endpoints = materials.NewTXMiddleware(srv, endpoints)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("test %s", <-c)
	}()

	go func() {
		_ = logger.Log("listening on port", *httpAddr)
		var serverOptions []httptransport.ServerOption
		handler := materials.NewHandler(ctx, serverOptions, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	_ = level.Error(logger).Log("exit", <-errs)
}
