package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var version string

type Flags struct {
	QueryKey       string
	FingerprintKey string
	SHA1           bool
	File           io.ReadCloser
}

func parseFlags() (flags *Flags) {
	flags = &Flags{}

	flag.StringVar(&flags.QueryKey, "q", "query", "query key")
	flag.StringVar(&flags.FingerprintKey, "f", "fingerprint", "fingerprint key")
	flag.BoolVar(&flags.SHA1, "sha1", false, "append SHA1")
	argVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *argVersion {
		printVersionAndEixt()
	}

	if flag.NArg() == 0 {
		flags.File = os.Stdin
	} else if flag.NArg() == 1 {
		file, err := os.OpenFile(flag.Arg(0), os.O_RDONLY, 0)

		if err != nil {
			log.Fatal(err)
		}

		flags.File = file
	} else {
		printUsageAndExit()
	}

	return
}

func printUsageAndExit() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func printVersionAndEixt() {
	fmt.Fprintln(os.Stderr, version)
	os.Exit(0)
}
