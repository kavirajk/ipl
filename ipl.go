package main

import (
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

	iplist := os.Args[1:]

	if len(iplist) == 0 {
		fmt.Println(help())
		os.Exit(1)
	}

	// sorry, API doesn't support bulk ips
	res := make([]ip.Info, 0, len(iplist))

	for _, ipadr := range iplist {
		r, err := ip.Lookup(ip.API, ipadr)
		if err != nil || r.Status != "success" {
			r.Err = errFailed
			r.IP = ipadr
		}
		res = append(res, r)
	}
	for _, i := range res {
		fmt.Println(ip.Format(i, verbose))
	}

}

func help() string {
	return fmt.Sprintf("Usage: ipl [IP-ADDRESS]...\nLook up IP-ADDRESSES.")
}
