package torrent

import (
	"github.com/anacrolix/torrent"
	"strings"
	"time"
)

func NewClient(config ClientConfig) ClientModel {
	c, err := torrent.NewClient(config.ClientConfig)
	if err != nil {
		panic("error on create new torrent client")
	}
	if config.TickerDuration.Seconds() == 0 {
		config.TickerDuration = 5 * time.Second
	}
	return &Client{Client: c, ticker: time.NewTicker(config.TickerDuration)}
}

func (c *Client) GetClient() *torrent.Client {
	return c.Client
}

func (c *Client) AddTorrent(s string) error {
	var t *torrent.Torrent
	var err error
	if strings.HasPrefix(s, "magnet") {
		t, err = c.AddMagnet(s)
	} else {
		t, err = c.AddTorrentFromFile(s)
	}
	if err != nil {
		return err
	}
	tor := NewTorrent(t, c.ticker)
	go tor.Initiate()
	c.torrents = append(c.torrents, tor)
	return nil
}

func (c *Client) GetTorrents() []TorrentModel {
	return c.torrents
}

func (c *Client) Start() {
	c.start = true
}

func (c *Client) Stop() {
	c.start = false
	c.Close()
}
