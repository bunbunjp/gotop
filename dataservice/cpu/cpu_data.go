package cpudata

import (
	"github.com/shirou/gopsutil/cpu"
	"time"
	"log"
	"strings"
	"runtime"
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
		if (runtime.GOOS == "windows") {
			c.update_windows()
		} else {
			c.update()
		}

		time.Sleep(1 * time.Second)
	}
}

func (c *Service) update_windows() {
	perf, _ := cpu.PerfInfo()

	idx := 0
	for _, core := range perf[2:] {
		name := strings.Replace(core.Name, " ", "", -1)

		log.Println("core is '", name, "', '", core.PercentProcessorTime, "'")
		c.AccumuData[idx] = append(c.AccumuData[idx], int(core.PercentProcessorTime))
		idx++
	}
}

func (c *Service) update() {
	percent, err := cpu.Percent(0*time.Millisecond, false)

	log.Println("percent, ", percent)
	log.Println("err, ", err)

	c.Latest = percent
	for idx, val := range c.Latest {
		if (len(c.AccumuData) >= idx) {
			break
		}
		c.AccumuData[idx] = append(c.AccumuData[idx], int(val))
	}

}
