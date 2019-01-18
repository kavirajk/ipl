package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/kavirajk/ipl/ip"
)

var (
	errFailed = errors.New("<failed>")
)

func main() {
	var verbose bool

	flag.BoolVar(&verbose, "v", false, "print verbose information")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	iplist := make([]string, 0)

	for scanner.Scan() {
		iplist = append(iplist, scanner.Text())
	}

	var (
		infoChan = make(chan ip.Info)
	)

	// Scatter
	for i := 0; i < len(iplist); i++ {
		ipadr := iplist[i]
		go func(ipadr string) {
			// sorry, API doesn't support bulk ips
			r, err := ip.Lookup(ip.API, ipadr)
			if err != nil || r.Status != "success" {
				r.Err = errFailed
			}
			infoChan <- r
		}(ipadr)
	}

	// Gatter
	for i := 0; i < len(iplist); i++ {
		fmt.Println(ip.Format(<-infoChan, verbose))
	}
}
