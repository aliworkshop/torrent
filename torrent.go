package torrent

import (
	"fmt"
	"github.com/anacrolix/torrent"
	"github.com/dustin/go-humanize"
	"log"
	"time"
)

func NewTorrent(tor *torrent.Torrent, ticker *time.Ticker) TorrentModel {
	return &Torrent{Torrent: tor, ticker: ticker}
}

// Initiate get torrent info and initialize torrent struct values
func (t *Torrent) Initiate() {
	<-t.Torrent.GotInfo()
	t.info = t.Torrent.Info()
	t.name = t.info.Name
	t.piecesNumber = t.info.NumPieces()
	t.piecesLength = t.info.PieceLength
	t.filesNumber = len(t.info.UpvertedFiles())
	t.totalLength = t.info.TotalLength()
	for _, file := range t.Files() {
		t.files = append(t.files, NewFile(file))
	}
	t.downloadedPieces, t.remainPieces, t.partialPieces = 0, 0, 0
	t.lastStats = t.GetStats()
}

// GetName returns torrent's name
func (t *Torrent) GetName() string {
	info := t.Torrent.Info()
	return info.Name
}

// GetPiecesNum returns torrent file pieces number
func (t *Torrent) GetPiecesNum() int {
	return t.piecesNumber
}

// GetPiecesLen returns torrent file pieces length
func (t *Torrent) GetPiecesLen() int64 {
	return t.piecesLength
}

// GetFilesNum returns torrent files number
func (t *Torrent) GetFilesNum() int {
	return t.filesNumber
}

// GetTorrentTotalLen returns torrent total length
func (t *Torrent) GetTorrentTotalLen() int64 {
	return t.totalLength
}

// GetFiles returns files of given torrent
func (t *Torrent) GetFiles() []FileModel {
	return t.files
}

// GetDownloaded returns downloaded pieces of torrent total pieces
func (t *Torrent) GetDownloaded() int {
	return t.downloadedPieces
}

// GetRemain returns remain pieces of torrent total pieces
func (t *Torrent) GetRemain() int {
	return t.remainPieces
}

// GetPartial returns partial pieces of torrent
func (t *Torrent) GetPartial() int {
	return t.partialPieces
}

// Download start given torrent downloading process
func (t *Torrent) Download() {
	go func() {
		for {
			<-t.ticker.C
			t.downloadedPieces, t.partialPieces = 0, 0
			psrs := t.PieceStateRuns()

			for _, r := range psrs {
				if r.Complete {
					t.downloadedPieces += r.Length
				}
				if r.Partial {
					t.partialPieces += r.Length
				}
			}
		}
	}()
	t.DownloadAll()
}

// DownloadLog return a simple downloading log just for fun ;)
// you can generate another download log instead
func (t *Torrent) DownloadLog() {
	for !t.Completed() {
		<-t.ticker.C

		line := fmt.Sprintf(
			"downloading %q: %s/%s, %d/%d pieces completed, (%d partial): %v/s\n",
			t.GetName(),
			humanize.Bytes(uint64(t.BytesCompleted())),
			humanize.Bytes(uint64(t.GetTorrentTotalLen())),
			t.GetDownloaded(),
			t.GetPiecesNum(),
			t.GetPartial(),
			t.DownloadRate(),
		)

		log.Println(line)
	}
}

// Completed return true if torrent file download is completed
func (t *Torrent) Completed() bool {
	return t.Complete.Bool()
}

// GetStats returns torrent file state
func (t *Torrent) GetStats() State {
	return State{
		state: t.Stats(),
		time:  time.Now(),
	}
}

// DownloadRate return download speed for given ticker duration in human-readable bytes
func (t *Torrent) DownloadRate() string {
	stats := t.Stats()
	byteRate := int64(time.Second)
	byteRate *= stats.BytesReadUsefulData.Int64() - t.lastStats.state.BytesReadUsefulData.Int64()
	byteRate /= int64(time.Now().Sub(t.lastStats.time))

	t.lastStats = State{time: time.Now(), state: stats}
	return humanize.Bytes(uint64(byteRate))
}
