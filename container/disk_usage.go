package container

import (
	"github.com/jroimartin/gocui"
)

// DiskUsageContainer ディスク使用量表示用コンテナー
type DiskUsageContainer struct {
	//gauge *termui.Gauge
}

// Initialize # Container Interface
func (d *DiskUsageContainer) Initialize() {
}

// UpdateData # Container Interface
func (d *DiskUsageContainer) UpdateData() {
}

// UpdateRender # Container Interface
func (d *DiskUsageContainer) UpdateRender() {
	//data := disk.GetInstance()
	//
	//d.gauge.Percent = int(data.UsageStat.UsedPercent)
	//d.gauge.BorderLabel = fmt.Sprintf("disk usage (%dGB / %dGB) %d％",
	//	int(util.Byte2GBi(data.UsageStat.Used)),
	//	int(util.Byte2GBi(data.UsageStat.Total)),
	//	int(data.UsageStat.UsedPercent))
}

// CreateUI # Container Interface
func (d *DiskUsageContainer) CreateUI(g *gocui.Gui) error {
	//d.gauge = termui.NewGauge()
	//d.gauge.Percent = 30
	//d.gauge.Width = termui.TermWidth() / 2
	//d.gauge.Height = 3
	//d.gauge.BorderLabel = fmt.Sprint("disk usage")
	//d.gauge.PercentColor = termui.ColorYellow
	//d.gauge.BarColor = termui.ColorMagenta
	//d.gauge.BorderFg = termui.ColorWhite
	//d.gauge.BorderLabelFg = termui.ColorMagenta
	//
	//return d.gauge

	return nil
}
