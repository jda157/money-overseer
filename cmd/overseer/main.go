package main

import (
	"flag"
	"fmt"
)

var (
	configFile = flag.String("config_file", "values/config.yaml", "path to configuration file")
)

func main() {
	fmt.Println("Hello Go!")
}
