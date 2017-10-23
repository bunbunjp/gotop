package disk

import (
	"github.com/shirou/gopsutil/disk"
	"time"
)

// Service ディスク使用量取得用データサービス
type Service struct {
	UsageStat *disk.UsageStat
}

var sharedInstance = &Service{}

// GetInstance get singleton instance
func GetInstance() *Service {
	return sharedInstance
}

// Initialize # DataService Interface
func (d *Service) Initialize() {
	go d.updateGoroutine()

}
func (d *Service) updateGoroutine() {
	for {
		d.update()

		time.Sleep(3 * time.Second)
	}
}
func (d *Service) update() {
	d.UsageStat, _ = disk.Usage("/")
}
