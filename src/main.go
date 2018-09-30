/* CLI startpoint - main package */
package main

import (
	"flag"
	"fmt"
)

/* Main tester function */
func main() {

	/* INPUT opts */
	optConfig := flag.String("config", "./config.json", "Config filename.")
	flag.Parse()

	configFile = optConfig

	fmt.Printf("#> Reading config from JSON file: %s\n", *optConfig)
	configParserFromFile(&config, optConfig)

	/* Make the tests */
	urlTestSetup(&config)
}
