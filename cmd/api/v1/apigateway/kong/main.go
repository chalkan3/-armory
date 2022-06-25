package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"scheduler/internal/api/v1/database/redis"
	listner "scheduler/internal/schedule/v1/listners/redis"

	"scheduler/pkg/etcd"
	"scheduler/pkg/scheduler"

	"github.com/go-kit/log"
)

var eventListeners = scheduler.Listeners{
	"create-redis": listner.Redis,
}

func main() {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var (
		httpPort    = fs.String("http_port", "9013", "application http port default 9000")
		_           = fs.String("grpc_port", "50053", "application grpc port default 50051")
		serviceName = fs.String("service_name", "group_service", "service name")
	)

	flag.Usage = fs.Usage
	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	database, err := etcd.NewETCD().Connection(2*time.Second, "127.0.0.1:2379")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	scheduler := scheduler.NewScheduler(eventListeners)
	scheduler.CheckEventsInInterval(ctx, 40*time.Second)

	logger := createLogger(*httpPort)
	repository := redis.NewRepository(database)

	routes := redis.NewHTTPServer(redis.NewService(repository, scheduler), logger)

	logger.Log(
		"service name", *serviceName,
		"msg", "HTTP",
		"addr", *httpPort,
		"prom-path", "/metrics")

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		for range interrupt {
			logger.Log("\n❌ Interrupt received closing...")
			cancel()
		}
	}()

	go func() {
		errs <- http.ListenAndServe(":"+*httpPort, routes)
	}()

	logger.Log("exit", <-errs)

}

func createLogger(port string) log.Logger {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stderr)
	logger = log.With(logger, "listen", port, "caller", log.DefaultCaller)
	return logger
}
