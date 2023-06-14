# Torrent Downloader

This is simple torrent downloader written in Go.
The purpose of this package is to download files using the BitTorrent protocol.
It provides interfaces to download and get torrent information.

## Features
* Download files using BitTorrent protocol
* Supports multiple concurrent downloads
* Supports torrent magnet url and file path as well
* Pause and resume downloads
* Track download process with download speed rates
* Manage multiple torrents at once

## Requirements
* Go 1.18 or later

## Installation
Install the package with `go get -v github.com/aliworkshop/torrent`

## Library examples
There are some small [examples](https://github.com/aliworkshop/torrent/tree/main/example) in the package documentation.

## Help
Communication about the project is primarily through [issue tracker](https://github.com/aliworkshop/torrent/issues).

#### torrent download

* Step 1: Create new client to add torrent files to.

`TickerDuration is timer duration to show download process `
```
  client := torrent.NewClient(torrent.ClientConfig{
    TickerDuration: 3 * time.Second,
  })
```

* Step 2: Add torrent files to client

```
    err := client.AddTorrent("magnet:?xt=urn:btih:AE204757FE376C70852CD5818B01870F05EE7064")

    if err != nil {
        log.Fatalln("error on add torrent magnet url: ", err.Error())
    }
```

* Step 3: Start torrent files concurrent download
```
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
```

And add `client.GetClient().WaitAll()` to wait for downloads to complete

## Contributor
[Ali Torabi](https://github.com/aliworkshop)
If this library helps you in anyway, show your love :heart: by putting a :star: on this project :v:
