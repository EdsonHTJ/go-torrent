package gTorrent

import (
	"fmt"
	"strings"
	"time"

	"github.com/cenkalti/rain/torrent"
)

type gTorrentCfg struct {
	DataDir string
}

var gtCfg *gTorrentCfg = &gTorrentCfg{}

func SetDataDir(dir string) {
	gtCfg.DataDir = dir
}

func LoadConfig() (torrent.Config, error) {
	cfg := torrent.DefaultConfig
	cfg.ResumeOnStartup = true
	if gtCfg.DataDir != "" {
		cfg.DataDir = gtCfg.DataDir
	}

	return cfg, nil
}

func GetSession() (*torrent.Session, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	return torrent.NewSession(cfg)
}

func DownloadTorrent(magneticLink string, acceptDuplicate bool) error {
	ses, err := GetSession()
	if err != nil {
		return err
	}
	defer ses.Close()

	magneticLink += ".resume"
	list := ses.ListTorrents()
	for _, t := range list {
		m, err := t.Magnet()
		if err != nil {
			return err
		}

		if MagneticComp(m, magneticLink) && !acceptDuplicate {
			return fmt.Errorf("duplicate torrent")
		}
	}

	_, err = ses.AddURI(magneticLink, nil)
	if err != nil {
		return err
	}

	return nil
}

func MagneticComp(magnetic1 string, magnetic2 string) bool {
	i := strings.Index(magnetic1, "&")
	return magnetic1[:i] == magnetic2[:i]
}

func ListTorrents() ([]*torrent.Torrent, error) {
	ses, err := GetSession()
	if err != nil {
		return nil, err
	}
	defer ses.Close()

	return ses.ListTorrents(), nil
}

func TorrentsStats(callback func(torrent.Torrent)) error {
	list, err := ListTorrents()
	if err != nil {
		return err
	}

	go func() {
		for {
			for _, t := range list {
				callback(*t)
			}

			time.Sleep(1 * time.Second)
		}
	}()

	return nil
}

func RemoveAllTorrents() error {
	ses, err := GetSession()
	if err != nil {
		return err
	}
	defer ses.Close()

	tlist := ses.ListTorrents()
	for _, t := range tlist {
		ses.RemoveTorrent(t.ID())
	}

	return nil
}

func isURI(arg string) bool {
	return strings.HasPrefix(arg, "magnet:") || strings.HasPrefix(arg, "http://") || strings.HasPrefix(arg, "https://")
}
