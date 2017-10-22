package memory

import (
	"github.com/bunbunjp/gotop/util"
	"github.com/shirou/gopsutil/mem"
	"time"
)

type MemorySwapDataService struct {
	SwapUsedHistory    []float64
	VirtualUsedHistory []float64
	LatestSwapStat     mem.SwapMemoryStat
	LatestVirtualStat  mem.VirtualMemoryStat
}

var sharedInstance *MemorySwapDataService = &MemorySwapDataService{}

func GetInstance() *MemorySwapDataService {
	return sharedInstance
}

func (m *MemorySwapDataService) Initialize() {
	vstat, _ := mem.VirtualMemory()
	m.VirtualUsedHistory = append(m.VirtualUsedHistory, util.Byte2GB(float64((*vstat).Total)))
	go m.updateGoroutine()
}

func (m *MemorySwapDataService) updateGoroutine() {
	for {
		m.update()

		time.Sleep(1 * time.Second)
	}
}

func (m *MemorySwapDataService) update() {
	sstat, _ := mem.SwapMemory()
	m.LatestSwapStat = *sstat

	vstat, _ := mem.VirtualMemory()
	m.LatestVirtualStat = *vstat

	m.SwapUsedHistory = append(m.SwapUsedHistory, util.Byte2GB(float64(m.LatestSwapStat.Used)))
	m.VirtualUsedHistory = append(m.VirtualUsedHistory, util.Byte2GB(float64(m.LatestVirtualStat.Used)))
}
