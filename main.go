package main

import (
	"flag"

	"github.com/kunikuni03/ciderblocker/cider"
)

func main() {
	var ciderParam string
	flag.StringVar(&ciderParam, "CIDER", "0.0.0.0/32", "Returns the available IP addresses.")

	cider.CheckFormat(ciderParam)
}
