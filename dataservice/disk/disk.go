package disk

import (
	"time"
	"github.com/shirou/gopsutil/disk"
)

type DiskDataService struct {
	UsageStat *disk.UsageStat
}

var sharedInstance *DiskDataService = &DiskDataService{}

func GetInstance() *DiskDataService {
	return sharedInstance
}

func (d *DiskDataService) Initialize() {
	go d.updateGoroutine()

}
func (d *DiskDataService) updateGoroutine() {
	for {
		d.update()

		time.Sleep(3 * time.Second)
	}
}
func (d *DiskDataService) update() {
	d.UsageStat, _ = disk.Usage("/")
}