package ip

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	cases := []struct {
		info    Info
		verbose bool
		exp     string
	}{
		{
			info: Info{Country: "A", City: "B", IP: "8.8.8.8", Lat: "1.0", Lon: "2.0"},
			exp:  "8.8.8.8         A B",
		},
		{
			info:    Info{Country: "A", City: "B", IP: "8.8.8.8", ISP: "C", Lat: "1.0", Lon: "2.0"},
			exp:     "8.8.8.8         A B C 1.0 2.0",
			verbose: true,
		},
	}

	for _, c := range cases {
		got := Format(c.info, c.verbose)
		assert.Equal(t, c.exp, got)
	}
}

func TestLookup(t *testing.T) {
	cases := []struct {
		ip, expCountryName string
	}{
		{"8.8.8.8", "US"},
		{"0.0.0.0", "AU"},
	}

	for _, c := range cases {
		info, err := Lookup(API, c.ip)
		assert.NoError(t, err)
		fmt.Printf("%+v\n", info)
		assert.Equal(t, c.expCountryName, info.CountryCode)
	}
}
