package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"PilotJobService/config"
	"PilotJobService/job"
	"PilotJobService/svc"
	"PilotJobService/utils/log"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	config.MustLoad(*configFile, &c)

	ctx := svc.MustNewServiceContext(&c)

	scheduler := job.NewScheduler(ctx)
	scheduler.Start()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGKILL)
	select {
	case sig := <-ch:
		log.Infof("Received signal %s, exiting...\n", sig)
		os.Exit(0)
	}
}
