package container

import (
	"github.com/jroimartin/gocui"
)

// VirtualMemoryUsageContainer 全体の仮想メモリの表示用コンテナ
type VirtualMemoryUsageContainer struct {
	//virtualGauge *termui.Gauge
}

// Initialize # Container Interface
func (v *VirtualMemoryUsageContainer) Initialize() {
}

// UpdateRender # Container Interface
func (v *VirtualMemoryUsageContainer) UpdateRender() {
	//data := memory.GetInstance()

	//v.virtualGauge.Percent = int(data.LatestVirtualStat.UsedPercent)
	//v.virtualGauge.BorderLabel = fmt.Sprintf("virtual usage (%.2fGB/%.2fGB)",
	//	util.Byte2GBi(data.LatestVirtualStat.Used),
	//	util.Byte2GBi(data.LatestVirtualStat.Total))
}

// CreateUI # Container Interface
func (v *VirtualMemoryUsageContainer) CreateUI(g *gocui.Gui) error {

	//v.virtualGauge = termui.NewGauge()
	//v.virtualGauge.Width = termui.TermWidth() / 4
	//v.virtualGauge.Height = 10
	//v.virtualGauge.LabelAlign = termui.AlignRight
	//
	//return v.virtualGauge[
	return nil
}
