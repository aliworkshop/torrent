package torrent

import (
	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
	"time"
)

type ClientConfig struct {
	*torrent.ClientConfig
	TickerDuration time.Duration
}

type Client struct {
	*torrent.Client
	ticker   *time.Ticker
	torrents []TorrentModel
	start    bool
}

type Torrent struct {
	*torrent.Torrent
	ticker           *time.Ticker
	info             *metainfo.Info
	name             string
	piecesNumber     int
	piecesLength     int64
	filesNumber      int
	totalLength      int64
	files            []FileModel
	downloadedPieces int
	remainPieces     int
	partialPieces    int
	lastStats        State
	tick             chan struct{}
}

type File struct {
	*torrent.File
	name   string
	path   string
	length int64
}

type State struct {
	state torrent.TorrentStats
	time  time.Time
}
