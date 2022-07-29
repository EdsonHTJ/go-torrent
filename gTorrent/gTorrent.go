package gTorrent

import (
	"log"
	"strings"
	"time"

	"github.com/cenkalti/rain/torrent"
)

func DownloadTorrent(dir string, magneticLink string) error {
	cfg := torrent.DefaultConfig
	cfg.DataDir = dir
	cfg.ResumeOnStartup = true
	ses, err := torrent.NewSession(cfg)
	if err != nil {
		return err
	}

	magneticLink += ".resume"
	tor, err := ses.AddURI(magneticLink, nil)
	if err != nil {
		return err
	}

	for range time.Tick(time.Second) {
		s := tor.Stats()
		log.Printf("Status: %s, Downloaded: %d, Total: %d,Peers: %d", s.Status.String(), s.Bytes.Completed, s.Bytes.Total, s.Peers.Total)
	}

	return nil
}

func ListTorrents() ([]*torrent.Torrent, error) {
	cfg := torrent.DefaultConfig
	cfg.ResumeOnStartup = true
	ses, err := torrent.NewSession(cfg)
	if err != nil {
		return nil, err
	}

	return ses.ListTorrents(), nil
}

func RemoveAllTorrents() error {
	cfg := torrent.DefaultConfig
	cfg.ResumeOnStartup = true
	ses, err := torrent.NewSession(cfg)
	if err != nil {
		return err
	}

	tlist := ses.ListTorrents()
	for _, t := range tlist {
		ses.RemoveTorrent(t.ID())
	}

	return nil
}

func isURI(arg string) bool {
	return strings.HasPrefix(arg, "magnet:") || strings.HasPrefix(arg, "http://") || strings.HasPrefix(arg, "https://")
}
