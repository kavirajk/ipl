# ipl - IP Lookup
Simple IP lookup tool built on top of [ip2country](https://ip2country.info) api.

## Install
```bash
go get -u -v github.com/kavirajk/ipl
```

## Usage
```bash
bash-4.3$ ipl 8.8.8.8
8.8.8.8         United-States

# multiple
bash-4.3$ ipl 123.125.71.77 13.66.139.0 195.154.122.121 195.154.123.109 207.46.13.12 207.46.13.176 88.7.217.130b
123.125.71.77   China
13.66.139.0     United-States
195.154.122.121 France
195.154.123.109 France
207.46.13.12    United-States
207.46.13.176   United-States
88.7.217.130b   <failed>

# ipv6
bash-4.3$ ipl 2001:4860:4860::8888
2001:4860:4860::8888 United-States

```
