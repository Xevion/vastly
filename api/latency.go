package api

import probing "github.com/prometheus-community/pro-bing"

type PingResult struct {
	Ip      string
	Latency float64
}

func GetLatency(ip string) (float64, error) {
	pinger, err := probing.NewPinger(ip)
	if err != nil {
		return 0, err
	}

	pinger
	return 0, nil
}
