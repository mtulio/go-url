package main

import (
	"flag"
	"fmt"
	"os"
)

/* Initialize */
func init() {
	/* Envs */
	/* INPUT opts */
	config.OptEmpty = true
	optConfig := flag.String("config", "", "Config filename. Eg.: -default=hack/config.json")
	optURL := flag.String("url", "", "URL to test.")
	optDNS := flag.Bool("dns", false, "Force resolve DNS and test each endpoint.")
	flag.Usage = usage
	flag.Parse()

	// -config is defined
	if *optConfig != "" {
		config.OptConfFile = *optConfig
		config.OptEmpty = false
	}

	// -url has more precedence than -config
	if *optURL != "" {
		config.OptURL = *optURL

		if config.OptConfFile != "" {
			fmt.Println("\n WARNING: you cannot use -conifg and -url at same time.")
			fmt.Println("\t Using -url option and ignoring -config parameter.")
		}
		config.OptEmpty = false
	} else if len(flag.Args()) == 1 {
		config.OptURL = flag.Arg(0)
		config.OptEmpty = false
	}

	// extra options
	if optDNS != nil {
		config.OptForceDNS = *optDNS
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [-dns] (-config=inputfile|-url=url|url)\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

/* Main tester function */
func main() {

	if (len(flag.Args()) < 1) && (config.OptEmpty) {
		fmt.Println("#> missing arguments")
		usage()
		os.Exit(1)
	}

	if config.OptURL == "" {
		fmt.Printf("#> Reading config from JSON file: %s\n", config.OptConfFile)
		configParserFromFile()
	} else {
		fmt.Printf("#> Reading config from Param\n")
		configParserFromParam()
	}

	/* Make the tests */
	urlTestLauncher()
}
