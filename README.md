# Transrss

[![Go Report Card](https://goreportcard.com/badge/github.com/grooveygr/transrss)](https://goreportcard.com/report/github.com/grooveygr/transrss)

Transrss glues torrent rss feeds to the Transmission bitorrent client


## Features

- Consume feeds as provided by [showrss.info](https://showrss.info)
- Rolling dedup cache of latest torrents
- (Very) simple cache persistence

## Building

```
git clone git@github.com:grooveygr/transrss.git
cd transrss
go build
```

Cross compile to any supported golang platform. example for RPI 3:
```
GOOS=linux GOARCH=arm GOARM=7 go build
```

## Usage


```
Usage of ./transrss:
  -cache string
        Cache path (default "cache.json")
  -cachesize int
        Maximum cache size (default 100)
  -feed string
        Feed URL
  -transmission string
        Full URL to transmission RPC (default "http://127.0.0.1:8181/transmission/rpc")
```

Automate rss checking using your favorite scheduler. Crontab example:

1. Create a `checkrss.sh` script for outputing logs to a file: 
```
#!/bin/bash
./transrss -feed [YOUR FEED URL] >> ~/rss.log 2>&1
```

2. Edit crontab file by running:
```
crontab -e
```

3. Add the relevant crontab entry (check every 30 minutes):
```
*/30 * * * * ~/checkrss.sh
```
