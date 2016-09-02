package main

import (
	"fmt"
	"os"

	"github.com/coralproject/scheduler/schedule"
	"github.com/gabelula/kit/cfg"
	"github.com/gabelula/kit/log"
)

// Config environmental variables.
const (
	cfgToken        = "TOKEN"
	cfgProject      = "PROJECT"
	cfgTarget       = "TARGET"
	cfgImage        = "IMAGE"
	cfgLoggingLevel = "LOGGING_LEVEL"
)

func main() {
	if err := cfg.Init(cfg.EnvProvider{Namespace: "SCHEDULE"}); err != nil {
		fmt.Println("Unable to initialize configuration")
		os.Exit(1)
	}

	logLevel := func() int {
		ll, err := cfg.Int(cfgLoggingLevel)
		if err != nil {
			return log.NONE
		}
		return ll
	}
	log.Init(os.Stderr, logLevel, log.Ldefault)

	// Pull options from the config.
	target := cfg.MustString(cfgTarget)
	image := cfg.MustString(cfgImage)

	schedule.Execute(target, image)
}
