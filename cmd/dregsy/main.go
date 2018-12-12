/*
 * TODO:
 *	- switch to log package
 *	-
 */

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xelalexv/dregsy/internal/pkg/log"
	"github.com/xelalexv/dregsy/internal/pkg/sync"
)

var DregsyVersion string

var syncInstance *sync.Sync
var dregsyExitCode int

var inTestRound bool

//
func version() {
	log.Info("\ndregsy %s\n", DregsyVersion)
}

//
func main() {

	dregsyExitCode = 0

	configFile := flag.String("config", "", "path to config file")
	flag.Parse()

	if len(*configFile) == 0 {
		version()
		fmt.Println("synopsis: dregsy -config={config file}")
		exit(1)
	}

	version()

	conf, err := sync.LoadConfig(*configFile)
	failOnError(err)

	syncInstance, err = sync.New(conf)
	failOnError(err)

	err = syncInstance.SyncFromConfig(conf)
	syncInstance.Dispose()
	failOnError(err)
}

//
func failOnError(err error) {
	if err != nil {
		log.Error(err)
		exit(1)
	}
}

//
func exit(code int) {
	dregsyExitCode = code
	if !inTestRound {
		os.Exit(code)
	}
}
