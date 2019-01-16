package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/kavirajk/ipl/ip"
)

const (
	api = "https://api.ip2country.info/ip?"
)

var (
	errFailed = errors.New("<failed>")
)

func main() {
	iplist := os.Args[1:]

	if len(iplist) == 0 {
		fmt.Println(help())
		os.Exit(1)
	}

	// sorry, API doesn't support bulk ips
	res := make([]ip.Info, 0, len(iplist))

	for _, ipadr := range iplist {
		r, err := ip.Lookup(api, ipadr)
		if err != nil || r.Country.Name == "" {
			r.Err = errFailed
			r.IP = ipadr
		}
		res = append(res, r)
	}
	fmt.Println(ip.Records(res))

}

func help() string {
	return fmt.Sprintf("Usage: ipl [IP-ADDRESS]...\nLook up IP-ADDRESSES.")
}
