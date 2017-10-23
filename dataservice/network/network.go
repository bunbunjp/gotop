package network

import (
	"github.com/shirou/gopsutil/net"
	"time"
)

// Service ネットワーク使用率のデータサービス
type Service struct {
	LatestStatus  net.IOCountersStat
	SentPerSecond uint64
	SentHistory   []int
	RecvPerSecond uint64
	RecvHistory   []int
}

var sharedInstance = &Service{}

// GetInstance is get singleton instance
func GetInstance() *Service {
	return sharedInstance
}

// Initialize is DataService interface
func (n *Service) Initialize() {
	n.SentHistory = []int{}

	n.RecvHistory = []int{}

	go n.updateGoroutine()
}

func (n *Service) updateGoroutine() {
	for {
		n.update()

		time.Sleep(1 * time.Second)
	}
}

func (n *Service) update() {
	statuses, _ := net.IOCounters(false)
	for _, s := range statuses {
		if s.Name == "all" {
			if n.LatestStatus.BytesRecv > 0 {
				n.SentPerSecond = s.BytesRecv - n.LatestStatus.BytesRecv
				n.RecvHistory = append(n.RecvHistory, int(n.SentPerSecond))
			}
			if n.LatestStatus.BytesSent > 0 {
				n.RecvPerSecond = s.BytesSent - n.LatestStatus.BytesSent
				n.SentHistory = append(n.SentHistory, int(n.RecvPerSecond))
			}

			n.LatestStatus = s

			break
		}
	}
}
