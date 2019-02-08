package main

import (
	"flag"
	"log"

	"github.com/Tubbebubbe/transmission"
	"github.com/mmcdole/gofeed"
)

func main() {
	fURL := flag.String("feed", "", "Feed URL")
	cachePath := flag.String("cache", "cache.json", "Cache path")
	cacheSize := flag.Int("cachesize", 100, "Maximum cache size")
	tRPC := flag.String("transmission", "http://127.0.0.1:8181/transmission/rpc", "Full URL to transmission RPC")

	flag.Parse()

	if *fURL == "" {
		log.Fatalln("Feed URL not provided")
	}

	log.Println("Running.")

	cache := newOrderedCache(*cachePath, *cacheSize)

	tclient := transmission.New(*tRPC, "", "")

	fparser := gofeed.NewParser()
	feed, err := fparser.ParseURL(*fURL)
	if err != nil {
		log.Fatalln("Failed parsing feed: ", err)
	}

	for _, item := range feed.Items {
		if cache.exists(item.Link) {
			continue
		}

		addcmd, err := transmission.NewAddCmdByMagnet(item.Link)
		if err != nil {
			log.Fatalln("Failed creating add cmd: ", err)
		}

		_, err = tclient.ExecuteAddCommand(addcmd)
		if err != nil {
			log.Fatalln("Failed adding torrent to transmission: ", err)
		}

		cache.add(item.Link)
		log.Println("Added: ", item.Title)
	}

	cache.commit()

	log.Println("Done.")
}
