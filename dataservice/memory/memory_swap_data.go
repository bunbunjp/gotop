package memory

import (
	"github.com/bunbunjp/gotop/util"
	"github.com/shirou/gopsutil/mem"
	"time"
)

// Service 仮想メモリ、スワップ領域などのデータサービスです
type Service struct {
	SwapUsedHistory    []float64
	VirtualUsedHistory []float64
	LatestSwapStat     mem.SwapMemoryStat
	LatestVirtualStat  mem.VirtualMemoryStat
}

var sharedInstance = &Service{}

// GetInstance is get singleton instance
func GetInstance() *Service {
	return sharedInstance
}

// Initialize is DataService interface
func (m *Service) Initialize() {
	vstat, _ := mem.VirtualMemory()
	m.VirtualUsedHistory = append(m.VirtualUsedHistory, util.Byte2GB(float64((*vstat).Total)))
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
	m.LatestSwapStat = *sstat

	vstat, _ := mem.VirtualMemory()
	m.LatestVirtualStat = *vstat

	m.SwapUsedHistory = append(m.SwapUsedHistory, util.Byte2GB(float64(m.LatestSwapStat.Used)))
	m.VirtualUsedHistory = append(m.VirtualUsedHistory, util.Byte2GB(float64(m.LatestVirtualStat.Used)))
}
