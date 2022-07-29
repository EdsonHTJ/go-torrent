package gTorrent_test

import (
	"go-torrent/gTorrent"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDownloadTorrent(t *testing.T) {
	err := gTorrent.DownloadTorrent("../data", "magnet:?xt=urn:btih:d67ed82392c28cb6c40509383ba70bfb4e6aefdf&tr=http://nyaa.tracker.wf:7777/announce&tr=udp://tracker.open-internet.nl:6969/announce&tr=udp://tracker.coppersurfer.tk:6969/announce&tr=https://1.track.ga:443/announce")
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
