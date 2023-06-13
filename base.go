package torrent

import (
	"github.com/anacrolix/torrent"
)

type ClientModel interface {
	GetClient() *torrent.Client
	AddTorrent(string) error
	GetTorrents() []TorrentModel
	Start()
	Stop()
}

type TorrentModel interface {
	Initiate()
	GetName() string
	GetPiecesNum() int
	GetPiecesLen() int64
	GetFilesNum() int
	GetTorrentTotalLen() int64
	GetFiles() []FileModel
	GetDownloaded() int
	GetRemain() int
	GetPartial() int
	Download()
	DownloadLog()
	Completed() bool
	GetStats() State
	DownloadRate() string
}

type FileModel interface {
	GetFile() *torrent.File
	GetName() string
	GetPath() string
	GetLength() int64
}
