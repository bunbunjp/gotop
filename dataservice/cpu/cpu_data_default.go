// +build linux darwin

package cpudata

import (
	"github.com/shirou/gopsutil/cpu"
	"time"
	"log"
)

func (c *Service) Update() {
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
