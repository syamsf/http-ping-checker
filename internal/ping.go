package internal

import (
	"github.com/go-ping/ping"
	"time"
)

func PingCmd(address string) error {
	pinger, err := ping.NewPinger(address)
	if err != nil {
		return err
	}

	pinger.Timeout = 15 * time.Second
	err = pinger.Run()
	if err != nil {
		return err
	}

	return nil
}
