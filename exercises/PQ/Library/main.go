package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-sql/sqlexp"
	"go-kit-template/config"
	"go-kit-template/materials"
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
	db, _ := sql.Open(sqlexp.DialectMySQL, config.Values.DB.DB_HOST)

	repository := materials.NewRepository(db)
	srv := materials.NewService(repository, logger)
	endpoints := materials.MakeEndpoints(srv)

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
