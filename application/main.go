package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	// exitFail is the exit code if the program fails.
	exitFail = 1
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	var (
		configPtr = flags.String("c", "config.json", "The configuration path 'it should be a json file'")
	)

	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	fmt.Fprintf(stdout, "Configuration path : %v\n", *configPtr)

	return nil
}

// TORUN:
// elwizarabot -c=config.json
