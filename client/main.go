package client

import (
	"bitbucket.org/kardianos/osext"
	"bitbucket.org/kardianos/service"
	"fmt"
	"github.com/inconshreveable/mousetrap"
	"github.com/thesyncim/at/log"
	"github.com/thesyncim/at/util"
	"math/rand"
	"os"

	"runtime"
	"time"
)

var logos service.Logger
var execdir, _ = osext.ExecutableFolder()

func init() {
	if runtime.GOOS == "windows" {
		if mousetrap.StartedByExplorer() {
			fmt.Println("Don't double-click ngrok!")
			fmt.Println("You need to open cmd.exe and run it from the command line!")
			time.Sleep(5 * time.Second)
			os.Exit(1)
		}
	}
}

func Main() {
	//os service
	var name = "anytunnel"
	var displayName = "anytunnel"
	var desc = "Introspected tunnels to localhost"
	var s service.Service

	s, err := service.NewService(name, displayName, desc)
	logos = s

	if err != nil {
		fmt.Printf("%s unable to start: %s", displayName, err)
		return
	}
	// parse options
	opts, err := ParseArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// set up logging
	log.LogTo(opts.logto)

	// read configuration file
	config, err := LoadConfiguration(opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// seed random number generator
	seed, err := util.RandomSeed()
	if err != nil {
		fmt.Printf("Couldn't securely seed the random number generator!")
		os.Exit(1)
	}
	rand.Seed(seed)

	switch opts.command {
	case "install":
		err = s.Install()
		if err != nil {
			fmt.Printf("Failed to install: %s\n", err)
			return
		}
		fmt.Printf("Service \"%s\" installed.\n", displayName)
	case "remove":
		err = s.Remove()
		if err != nil {
			fmt.Printf("Failed to remove: %s\n", err)
			return
		}
		fmt.Printf("Service \"%s\" removed.\n", displayName)
	case "run":
		doWork(config)

		//continue
	case "start":
		err = s.Start()
		if err != nil {
			fmt.Printf("Failed to start: %s\n", err)
			return
		}
		fmt.Printf("Service \"%s\" started.\n", displayName)
	case "stop":
		err = s.Stop()
		if err != nil {
			fmt.Printf("Failed to stop: %s\n", err)
			return
		}
		fmt.Printf("Service \"%s\" stopped.\n", displayName)

	}
	err = s.Run(func() error {
		// start
		go doWork(config)
		return nil
	}, func() error {
		os.Exit(0)
		// stop
		//stopWork()
		return nil
	})
	if err != nil {
		s.Error(err.Error())
	}

}

func doWork(c *Configuration) {
	NewController().Run(c)

}
func stopWork() {

}
