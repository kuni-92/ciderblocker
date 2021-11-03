package main

import (
	"flag"
	"fmt"

	"github.com/kunikuni03/ciderblocker/cider"
)

func main() {
	var ciderParam string
	flag.StringVar(&ciderParam, "CIDER", "0.0.0.0/32", "Returns the available IP addresses.")

	_, _, err := cider.CheckFormat(ciderParam)
	if err != nil {
		fmt.Println("CIDER is invalid format.")
		fmt.Println(err)
	}

}
