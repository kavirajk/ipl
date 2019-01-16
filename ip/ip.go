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

type Country struct {
	Code  string `json:"countryCode"`
	Code3 string `json:"countryCode3"`
	Name  string `json:"countryName"`
	Emoji string `json:"countryEmoji"`
}

func (c Country) String() string {
	return fmt.Sprintf("%s", c.Name)
}

type Info struct {
	Country Country
	IP      string
	Err     error
}

func (i Info) String() string {
	if i.Err == nil {
		return fmt.Sprintf("%-15s %s", i.IP, i.Country)
	}
	return fmt.Sprintf("%-15s %s", i.IP, i.Err)
}

type Records []Info

func (rs Records) String() string {
	res := make([]string, 0, len(rs))
	for _, r := range rs {
		res = append(res, r.String())
	}
	return strings.Join(res, "\n")
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

	var country Country

	if err := json.Unmarshal(body, &country); err != nil {
		return Info{}, errors.Wrap(err, "failed to unmarshal response")
	}
	country.Name = strings.Join(strings.Split(country.Name, " "), "-")
	return Info{Country: country, IP: ip}, nil
}
