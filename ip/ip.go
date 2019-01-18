package ip

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	API = "https://extreme-ip-lookup.com/json/"
)

type Info struct {
	BusinessName    string `json:"businessName"`
	BusinessWebsite string `json:"businessWebsite"`
	City            string `json:"city"`
	Continent       string `json:"continent"`
	Country         string `json:"country"`
	CountryCode     string `json:"countryCode"`
	IPName          string `json:"ipName"`
	IPType          string `json:"ipType"`
	ISP             string `json:"isp"`
	Lat             string `json:"lat"`
	Lon             string `json:"lon"`
	Org             string `json:"org"`
	IP              string `json:"query"`
	Region          string `json:"region"`
	Status          string `json:"status"`
	Code            string `json:"countryCode"`
	Code3           string `json:"countryCode3"`
	Name            string `json:"countryName"`
	Emoji           string `json:"countryEmoji"`
	Err             error  `json:"-"`
}

func Format(i Info, verbose bool) string {
	ip := fmt.Sprintf("%-39s", i.IP) // max upto ipv6

	if i.Err != nil {
		return fmt.Sprintf("%s %v", ip, i.Err)
	}
	if !verbose {
		return fmt.Sprintf("%s %-20s %-10s", ip, normalize(i.Country), normalize(i.City))
	}
	// ip, country, city, isp
	return fmt.Sprintf("%s %-20s %-10s %-30s", ip, normalize(i.Country), normalize(i.City), normalize(i.ISP))
}

func normalize(s string) string {
	if s == "" {
		return "<NA>"
	}
	return s
}

type Error struct {
	body []byte
	r    *http.Response
}

func (e *Error) Error() string {
	return fmt.Sprintf("Status: %s\n%s", e.r.Status, e.body)
}

func Lookup(api, ip string) (Info, error) {
	c := http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("GET", api+ip, nil)
	if err != nil {
		return Info{}, errors.Wrap(err, "failed to create api request")
	}
	resp, err := c.Do(req)
	if err != nil {
		return Info{}, errors.Wrap(err, "failed to make request to api")
	}

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return Info{}, errors.Wrap(err, "failed to read response body")
	}

	if code := resp.StatusCode; code < 200 || code > 299 {
		return Info{}, &Error{r: resp, body: body}
	}

	var info Info

	if err := json.Unmarshal(body, &info); err != nil {
		return Info{}, errors.Wrap(err, "failed to unmarshal response")
	}
	info.Name = strings.Join(strings.Split(info.Name, " "), "-")
	info.Continent = strings.Join(strings.Split(info.Continent, " "), "-")
	info.Country = strings.Join(strings.Split(info.Country, " "), "-")
	info.City = strings.Join(strings.Split(info.City, " "), "-")
	info.ISP = strings.Join(strings.Split(info.ISP, " "), "-")

	return info, nil
}
