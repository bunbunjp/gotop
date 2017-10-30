// +build windows

package cpudata

import (
	"github.com/shirou/gopsutil/cpu"
	"strings"
	"log"
)

func (c *Service) Update() {
	perf, _ := cpu.PerfInfo()

	idx := 0
	for _, core := range perf[2:] {
		name := strings.Replace(core.Name, " ", "", -1)

		log.Println("core is '", name, "', '", core.PercentProcessorTime, "'")
		c.AccumuData[idx] = append(c.AccumuData[idx], int(core.PercentProcessorTime))
		idx++
	}
}
