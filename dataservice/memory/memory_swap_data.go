package memory

import (
	"github.com/shirou/gopsutil/mem"
	"time"
)

// Service 仮想メモリ、スワップ領域などのデータサービスです
type Service struct {
	SwapUsedHistory    []uint64
	VirtualUsedHistory []uint64
	LatestSwapStat     mem.SwapMemoryStat
	LatestVirtualStat  mem.VirtualMemoryStat
	TotalMemorySize    uint64
}

var sharedInstance = &Service{}

// GetInstance is get singleton instance
func GetInstance() *Service {
	return sharedInstance
}

// Initialize is DataService interface
func (m *Service) Initialize() {
	vstat, _ := mem.VirtualMemory()
	m.TotalMemorySize = (*vstat).Total
	go m.updateGoroutine()
}

func (m *Service) updateGoroutine() {
	for {
		m.update()

		time.Sleep(1 * time.Second)
	}
}

func (m *Service) update() {
	sstat, _ := mem.SwapMemory()
	m.TotalMemorySize = sstat.Total
	m.LatestSwapStat = *sstat

	vstat, _ := mem.VirtualMemory()
	m.LatestVirtualStat = *vstat

	m.SwapUsedHistory = append(m.SwapUsedHistory, m.LatestSwapStat.Used)
	m.VirtualUsedHistory = append(m.VirtualUsedHistory, m.LatestVirtualStat.Used)
}
