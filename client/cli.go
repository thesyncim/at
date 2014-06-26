package client

import (
	"flag"
	"fmt"
	"path/filepath"
	//"github.com/thesyncim/at/version"
	"os"
)

const usage1 string = `Usage: %s [OPTIONS] <local port or address>
Options:
`

const usage2 string = `
Examples:
	at -config config.yaml
`

type Options struct {
	config    string
	logto     string
	authtoken string
	httpauth  string
	hostname  string
	protocol  string
	subdomain string
	command   string
	args      []string
}

func ParseArgs() (opts *Options, err error) {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage1, os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, usage2)
	}

	config := flag.String(
		"config", filepath.Join(execdir+"config.yaml"),
		"Path to ngrok configuration file. (default: $HOME/.ngrok)")

	logto := flag.String(
		"log",
		"none",
		"Write log messages to this file. 'stdout' and 'none' have special meanings")

	authtoken := flag.String(
		"authtoken",
		"",
		"Authentication token for identifying an ngrok.com account")

	flag.Parse()

	opts = &Options{
		config:    *config,
		logto:     *logto,
		authtoken: *authtoken,
		command:   flag.Arg(0),
	}

	/*switch opts.command {
	case "start":
		opts.args = flag.Args()[1:]
	case "version":
		fmt.Println(version.MajorMinor())
		os.Exit(0)
	case "help":
		flag.Usage()
		os.Exit(0)
	case "":
		if opts.config == "" {
			err = fmt.Errorf("Error: Specify a local port to tunnel to, or " +
				"an AnnyTunnel command.\n\nExample: To expose port 80, run " +
				"'at 80'")

		}

		return

	default:
		if len(flag.Args()) > 1 {
			err = fmt.Errorf("You may only specify one port to tunnel to on the command line, got %d: %v",
				len(flag.Args()),
				flag.Args())
			return
		}

		opts.command = "default"
		opts.args = flag.Args()
	}*/

	return
}
