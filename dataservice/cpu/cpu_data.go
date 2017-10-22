package cpudata


import (
	"github.com/shirou/gopsutil/cpu"
	"time"
)

type CpuDataService struct  {
	CoreCount int
	Latest []float64
	AccumuData [][]int

	isRunning bool
}

var sharedInstance *CpuDataService = &CpuDataService{isRunning:false}

func GetInstance() *CpuDataService  {
	return sharedInstance
}

func (c *CpuDataService) Initialize() {
	if c.isRunning {
		return
	}
	core, _ := cpu.Counts(true)
	c.CoreCount = core
	c.AccumuData = make([][]int, core, core)
	for i:=0; i<c.CoreCount; i++ {
		c.AccumuData[i] = append(c.AccumuData[i], 100)
	}

	go c.updateGoroutine()
	c.isRunning = true
}

func (c *CpuDataService) updateGoroutine() {
	for {
		c.update()

		time.Sleep(1 * time.Second)
	}
}

func (c *CpuDataService) update() {
	percent, _ := cpu.Percent(0 * time.Millisecond, true)

	c.Latest = percent
	for idx, val := range c.Latest {
		c.AccumuData[idx] = append(c.AccumuData[idx], int(val))
	}

}