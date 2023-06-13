package main

import (
	"context"
	"github.com/aliworkshop/torrent"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

func main() {
	client := torrent.NewClient(torrent.ClientConfig{
		TickerDuration: 3 * time.Second,
	})
	defer client.Stop()
	err := client.AddTorrent("magnet:?xt=urn:btih:AE204757FE376C70852CD5818B01870F05EE7064")
	if err != nil {
		log.Fatalln("error on add torrent magnet url")
	}

	eg, _ := errgroup.WithContext(context.Background())
	for _, tt := range client.GetTorrents() {
		eg.Go(func(t torrent.TorrentModel) func() error {
			return func() error {
				t.Initiate()
				t.Download()

				go t.DownloadLog()
				return nil
			}
		}(tt))
	}
	eg.Wait()

	client.GetClient().WaitAll()
	log.Print("congratulations, all torrents downloaded!")
}
