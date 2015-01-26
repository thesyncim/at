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
		"Path to AnyTunnel configuration file. (default: cinfig.yaml)")

	logto := flag.String(
		"log",
		"none",
		"Write log messages to this file. 'stdout' and 'none' have special meanings")

	authtoken := flag.String(
		"authtoken",
		"",
		"Authentication token for identifying an proxy.alpeca3d.com account")

	flag.Parse()

	opts = &Options{
		config:    *config,
		logto:     *logto,
		authtoken: *authtoken,
		command:   flag.Arg(0),
	}
	return
}
