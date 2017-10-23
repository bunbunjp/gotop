package cpudata

import (
	"github.com/shirou/gopsutil/cpu"
	"time"
)

// Service CPU使用率のデータサービス
type Service struct {
	CoreCount  int
	Latest     []float64
	AccumuData [][]int

	isRunning bool
}

var sharedInstance = &Service{isRunning: false}

// GetInstance get singleton instance
func GetInstance() *Service {
	return sharedInstance
}

// Initialize # DataService Interface
func (c *Service) Initialize() {
	if c.isRunning {
		return
	}
	core, _ := cpu.Counts(true)
	c.CoreCount = core
	c.AccumuData = make([][]int, core, core)
	for i := 0; i < c.CoreCount; i++ {
		c.AccumuData[i] = append(c.AccumuData[i], 100)
	}

	go c.updateGoroutine()
	c.isRunning = true
}

func (c *Service) updateGoroutine() {
	for {
		c.update()

		time.Sleep(1 * time.Second)
	}
}

func (c *Service) update() {
	percent, _ := cpu.Percent(0*time.Millisecond, true)

	c.Latest = percent
	for idx, val := range c.Latest {
		c.AccumuData[idx] = append(c.AccumuData[idx], int(val))
	}

}
