package network

import (
	"time"
	"github.com/shirou/gopsutil/net"
)

type NetworkDataService struct {
	LatestStatus net.IOCountersStat
	SentPerSecond uint64
	SentHistory []int
	RecvPerSecond uint64
	RecvHistory []int
}

var sharedInstance *NetworkDataService = &NetworkDataService{
}

func GetInstace() *NetworkDataService {
	return sharedInstance
}

func (n *NetworkDataService) Initialize() {
	n.SentHistory = []int{}

	n.RecvHistory = []int{}

	go n.updateGoroutine()
}

func (n *NetworkDataService) updateGoroutine() {
	for {
		n.update()

		time.Sleep(1 * time.Second)
	}
}

func (n *NetworkDataService) update() {
	statuses, _ := net.IOCounters(false)
	for _, s := range statuses {
		if s.Name == "all" {
			if (n.LatestStatus.BytesRecv > 0) {
				n.SentPerSecond = s.BytesRecv - n.LatestStatus.BytesRecv
				n.RecvHistory = append(n.RecvHistory, int(n.SentPerSecond))
			}
			if (n.LatestStatus.BytesSent > 0) {
				n.RecvPerSecond = s.BytesSent - n.LatestStatus.BytesSent
				n.SentHistory = append(n.SentHistory, int(n.RecvPerSecond))
			}

			n.LatestStatus = s

			break
		}
	}
}