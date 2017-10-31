// +build linux darwin

package cpudata

import (
	"github.com/shirou/gopsutil/cpu"
	"time"
)

func (c *Service) Update() {
	percent, _ := cpu.Percent(0*time.Millisecond, false)

	c.Latest = percent
	for idx, val := range c.Latest {
		if len(c.AccumuData) >= idx {
			break
		}
		c.AccumuData[idx] = append(c.AccumuData[idx], int(val))
	}

}
