package api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
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
	ipTicker       *time.Ticker
	ipSelf         net.IP
	handlerChannel chan<- PingResult
}

func NewLatencyQueue(redis *redis.Client) *LatencyQueue {
	logger, _ := zap.NewDevelopment()

	return &LatencyQueue{
		processChannel: make(chan LatencyRequest, 1024),
		logger:         logger.Sugar(),
		redis:          redis,
		ipTicker:       time.NewTicker(time.Minute * 5),
	}
}

type PingRequest struct {
	Ip net.IPAddr
}

type PingResult struct {
	Ip         net.IPAddr
	Latency    time.Duration
	Successful bool
}

func (l *LatencyQueue) QueuePing(ip string) error {
	// Parse the IP
	parsedIp := net.ParseIP(ip)
	if parsedIp == nil {
		return errors.New("Invalid IP address")
	}

	// Create the request
	request := LatencyRequest{
		RequestTime: time.Now().UnixMilli(),
		Ip:          net.IPAddr{IP: parsedIp},
	}

	// Add the request to the queue
	l.processChannel <- request

	return nil
}

func (l *LatencyQueue) SetHandler(handler chan<- PingResult) {
	l.handlerChannel = handler
}

func (l *LatencyQueue) GetSelfIP() net.IP {
	return l.ipSelf
}

func (l *LatencyQueue) RefreshIP() error {
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		l.logger.Errorw("Failed to get IP address", "error", err)
		return err
	}
	defer resp.Body.Close()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		l.logger.Errorw("Failed to read response body", "error", err)
		return err
	}

	parsedIp := net.ParseIP(string(ip))
	if parsedIp == nil {
		l.logger.Errorw("Invalid IP address", "ip", ip)
		return errors.New("Invalid IP address")
	}

	l.logger.Debugw("IP Address Refreshed", "ip", parsedIp)

	l.ipSelf = parsedIp
	return nil
}

func (l *LatencyQueue) Start(ctx context.Context) {
	l.RefreshIP()
	for {
		select {
		case <-l.ipTicker.C:
			l.RefreshIP()
		case request := <-l.processChannel:
			ip := request.Ip.String()

			// Check if we have a result in the cache
			latencyKey := fmt.Sprintf("latency:%s:%s", l.ipSelf, ip)
			existingResult, err := l.redis.Get(ctx, latencyKey).Result()
			if err != nil {
				if err != redis.Nil {
					l.logger.Errorw("Failed to get existing result", "key", latencyKey, "ip", ip, "error", err)
					continue
				}
			} else {
				// Emit the result
				if (l.handlerChannel) != nil {
					parsed, err := time.ParseDuration(existingResult)
					if err != nil {
						l.logger.Errorw("Failed to parse existing result", "key", latencyKey, "ip", ip, "error", err)
						continue
					}

					l.handlerChannel <- PingResult{
						Ip:      request.Ip,
						Latency: parsed,
					}
				}
				continue
			}

			pinger, err := probing.NewPinger(ip)
			if err != nil {
				l.logger.Errorf("Failed to create pinger for %s: %s", ip, err)
				continue
			}

			pinger.SetPrivileged(true)
			pinger.Count = 1
			pinger.Interval = time.Nanosecond
			pinger.Timeout = time.Millisecond * 500

			// Process the request
			err = pinger.Run()
			if err != nil {
				l.logger.Errorf("Failed to ping %s: %s", ip, err)
				continue
			}

			// Get the results
			results := pinger.Statistics()
			success := results.PacketLoss == 0

			// Store the result in Redis
			value := "timeout"
			expiration := time.Hour*24 + time.Minute*time.Duration(rand.Intn(60*8))
			if success {
				value = results.AvgRtt.String()
			}

			l.logger.Debugw("Ping Result", "ip", ip, "latency", value)
			l.redis.SetEx(ctx, latencyKey, value, expiration)

			// Emit the result
			if (l.handlerChannel) != nil {
				l.handlerChannel <- PingResult{
					Ip:         request.Ip,
					Latency:    results.AvgRtt,
					Successful: success,
				}
			}

		case <-l.stopChannel:
			return
		}
	}
}

func (l *LatencyQueue) Kill() error {
	l.logger.Warn("Killing LatencyQueue")
	l.stopChannel <- true
	err := l.redis.Close()
	l.ipTicker.Stop()
	l.logger.Sync()
	return err
}
