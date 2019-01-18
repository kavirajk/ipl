# ipl - IP Lookup
Simple IP lookup tool built on top of [extreme-ip-lookup](https://extreme-ip-lookup.com) api.
Reads input from STDIN.

## Install
```bash
go get -u -v github.com/kavirajk/ipl
```

## Usage
```bash
$ echo "8.8.8.8" | ipl
8.8.8.8         United-States

# multiple
$ echo "123.125.71.77 13.66.139.0" | ipl
123.125.71.77   China
13.66.139.0     United-States

# ipv6
$ echo "2001:4860:4860::8888" | ipl
2001:4860:4860::8888 United-States

# from files
$ cat ips.txt | ipl
123.125.71.77   China
13.66.139.0     United-States
195.154.122.121 France
195.154.123.109 France
207.46.13.12    United-States
207.46.13.176   United-States

# verbose
$ cat ips.txt | ipl -v
13.66.139.0                             United-States        <NA>       Microsoft-Azure
195.154.123.109                         France               <NA>       Iliad-Entreprises
207.46.13.176                           United-States        <NA>       Microsoft-bingbot
123.125.71.77                           China                Beijing    China-Unicom-Beijing
207.46.13.12                            United-States        <NA>       Microsoft-bingbot
195.154.122.121                         France               <NA>       Iliad-Entreprises

```

## LICENSE
MIT License

Copyright (c) 2019 Kaviraj
