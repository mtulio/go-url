/* CLI startpoint - main package */
package main

import (
	"flag"
	"fmt"
)

/* Initialize */
func init() {
	/* Envs */
	/* INPUT opts */
	optConfig := flag.String("config", "./config.json", "Config filename.")
	optDNS := flag.Bool("dns", false, "Force resolve DNS and test each endpoint.")
	flag.Parse()

	if optConfig != nil {
		config.OptConfFile = *optConfig
	}
	if optDNS != nil {
		config.OptForceDNS = *optDNS
	}
}

/* Main tester function */
func main() {

	fmt.Printf("#> Reading config from JSON file: %s\n", config.OptConfFile)
	configParserFromFile()

	/* Make the tests */
	urlTestLauncher()
}
