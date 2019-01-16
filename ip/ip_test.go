package ip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	api = "https://api.ip2country.info/ip?"
)

func TestLookup(t *testing.T) {
	cases := []struct {
		ip, expCountryName string
	}{
		{"8.8.8.8", "USA"},
		{"0.0.0.0", "AUS"},
	}

	for _, c := range cases {
		info, err := Lookup(api, c.ip)
		assert.NoError(t, err)
		assert.Equal(t, c.expCountryName, info.Country.Code3)
	}
}
