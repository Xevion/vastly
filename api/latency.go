package api

import (
	"errors"
	"net"
	"time"

	probing "github.com/prometheus-community/pro-bing"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type LatencyRequest struct {
	RequestTime int64
	Ip          net.IPAddr
}

type LatencyQueue struct {
	processChannel chan LatencyRequest
	stopChannel    chan bool
	logger         *zap.SugaredLogger
	redis          *redis.Client
	pinger         *probing.Pinger
	handlerChannel chan<- PingResult
}

func NewLatencyQueue() *LatencyQueue {
	logger, _ := zap.NewDevelopment()
	pinger, err := probing.NewPinger("127.0.0.1")
	if err != nil {
		logger.Fatal("Failed to create pinger")
	}
	pinger.Count = 1
	pinger.Interval = time.Millisecond * 850
	return &LatencyQueue{
		processChannel: make(chan LatencyRequest, 1024),
		logger:         logger.Sugar(),
		redis:          redis.NewClient(&redis.Options{}),
		pinger:         pinger,
	}
}

type PingRequest struct {
	Ip net.IPAddr
}

type PingResult struct {
	Ip      net.IPAddr
	Latency int64
}

func (l *LatencyQueue) QueuePing(ip string) error {
	// Parse the IP
	parsedIp := net.ParseIP(ip)
	if parsedIp == nil {
		return errors.New("Invalid IP address")
	}

	// Create the request
	request := LatencyRequest{
		RequestTime: time.Now().Unix(),
		Ip:          net.IPAddr{IP: parsedIp},
	}

	// Add the request to the queue
	l.processChannel <- request

	return nil
}

func (l *LatencyQueue) Start() {
	for {
		select {
		case request := <-l.processChannel:

			ip := request.Ip.String()
			l.pinger.SetIPAddr(&request.Ip)

			// Process the request
			err := l.pinger.Run()
			if err != nil {
				l.logger.Errorf("Failed to ping %s: %s", ip, err)
				continue
			}

			// Get the results
			results := l.pinger.Statistics()
			if (l.handlerChannel) != nil {
				l.handlerChannel <- PingResult{
					Ip:      request.Ip,
					Latency: results.AvgRtt.Milliseconds(),
				}
			}

		case <-l.stopChannel:
			return
		}
	}
}

func (l *LatencyQueue) Kill() error {
	l.stopChannel <- true
	err := l.redis.Close()
	return err
}
