package svc

import (
	"fmt"
	"os"

	"PilotJobService/config"
	"PilotJobService/dal"
	"PilotJobService/utils/log"
)

type ServiceContext struct {
	SvcConf *config.Config
}

func MustNewServiceContext(c *config.Config) *ServiceContext {
	var (
		err error
	)
	err = log.InitLogger(c.AppName, c.LogPath, c.DebugMode)
	if err != nil {
		panic(fmt.Errorf("failed to create logger: %w", err))
	}

	if os.Getenv("MYSQL_DSN") != "" {
		c.MySQL = os.Getenv("MYSQL_DSN")
	}
	err = dal.InitDB(c.MySQL)
	if err != nil {
		panic(fmt.Errorf("failed to init db: %w", err))
	}
	return &ServiceContext{
		SvcConf: c,
	}
}
