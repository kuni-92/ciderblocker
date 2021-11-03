package main

import (
	"flag"
	"fmt"

	"github.com/kunikuni03/ciderblocker/cider"
)

func main() {
	var ciderParam string
	flag.StringVar(&ciderParam, "CIDER", "0.0.0.0/32", "Returns the available IP addresses.")
	flag.Parse()

	addr, mask, err := cider.CheckFormat(ciderParam)
	if err != nil {
		fmt.Println("CIDER is invalid format.")
		fmt.Println(err)
	}

	subnet := cider.GetSubnetmask(mask)

	naddr := cider.GetNetworkAddress(addr, subnet)
	baddr := cider.GetBroadcastAddress(addr, subnet)

	msg := "address range: %d.%d.%d.%d - %d.%d.%d.%d\n"
	fmt.Printf(msg, naddr[0], naddr[1], naddr[2], naddr[3], baddr[0], baddr[1], baddr[2], baddr[3])

}
