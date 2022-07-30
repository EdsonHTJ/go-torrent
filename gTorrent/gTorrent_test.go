package gTorrent_test

import (
	"fmt"
	"go-torrent/gTorrent"
	"testing"
	"time"

	"github.com/cenkalti/rain/torrent"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	gTorrent.SetDataDir("../data")
	m.Run()
}

func TestDownloadTorrent(t *testing.T) {
	err := gTorrent.DownloadTorrent("magnet:?xt=urn:btih:d67ed82392c28cb6c40509383ba70bfb4e6aefdf&tr=http://nyaa.tracker.wf:7777/announce&tr=udp://tracker.open-internet.nl:6969/announce&tr=udp://tracker.coppersurfer.tk:6969/announce&tr=https://1.track.ga:443/announce", false)
	require.NoError(t, err)
}

func TestListTorrent(t *testing.T) {
	tlist, err := gTorrent.ListTorrents()
	require.NoError(t, err)
	t.Log(tlist)
}

func TestRemoveTorrents(t *testing.T) {
	err := gTorrent.RemoveAllTorrents()
	require.NoError(t, err)
}

func TestTorrentStats(t *testing.T) {
	err := gTorrent.TorrentsStats(func(t torrent.Torrent) {
		fmt.Printf("downloaded: %d, total: %d\n", t.Stats().Bytes.Completed, t.Stats().Bytes.Total)
	})
	time.Sleep(20 * time.Second)
	require.NoError(t, err)
}
