package main

import (
	"context"
	"flag"
	"github.com/4thel00z/libhttp"
	"github.com/joho/godotenv"
	"github.com/logrusorgru/aurora"

	libservice "github.com/4thel00z/libservice/v1"
	"log"
	"os"
	"os/signal"
	"service_templated/pkg/libservice_template/filters"
	"service_templated/pkg/libservice_template/modules/debug"
	"strconv"
	"syscall"
	"time"
)

const (
	banner = `
SERVICE_TEMPLATE_BANNER
	`
	version = "0.0.1"
)

var (
	port       = flag.Int("port", 1337, "<port> [defaults to 1337]")
	host       = flag.String("host", "0.0.0.0", "<host> [defaults to 0.0.0.0]")
	configPath = flag.String("config", ".service_templatedrc.json", "path to config. [defaults to .service_templatedrc.json]")
	debugFlag  = flag.Bool("debug", false, "enable replacement of <auth> with real token [defaults to false]")
	verbose    = flag.Bool("verbose", false, "enable verbose logging [defaults to false]")
)

func main() {
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Create an empty .env file if you don't intend to use them")
	}

	log.Println("\n", aurora.Magenta(banner))
	log.Println("\n", "👩	Version:", version)

	config, err := libservice.ParseConfig(*configPath)
	if err != nil {
		log.Fatalf("could not parse the configuration because of: %s", err.Error())
	}
	addr := *host + ":" + strconv.Itoa(*port)
	app := libservice.NewApp(addr, config, *verbose, *debugFlag, debug.Module)

	svc := app.Router.Serve().
		Filter(libhttp.ErrorFilter).
		Filter(libhttp.H2cFilter).
		Filter(filters.Validation(app)).
		Filter(filters.MultipartValidation(app)).
		Filter(filters.Auth(app))

	srv, err := libhttp.Listen(svc, addr)
	if err != nil {
		panic(err)
	}

	log.Printf("🏁	Listening on %v", srv.Listener().Addr())
	if app.Debug {
		app.PrintConfig()
	}
	app.PrintRoutes(srv.Listener().Addr().String())
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.Printf("☠️  Shutting down in max 10 sec..")
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Stop(c)
}
