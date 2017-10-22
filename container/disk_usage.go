package container

import (
	"fmt"
	"github.com/bunbunjp/gotop/dataservice/disk"
	"github.com/bunbunjp/gotop/util"
	"github.com/gizak/termui"
)

type DiskUsageContainer struct {
	gauge *termui.Gauge
}

func (d *DiskUsageContainer) Initialize() {
}

func (d *DiskUsageContainer) UpdateData() {
}

func (d *DiskUsageContainer) UpdateRender() {
	data := disk.GetInstance()

	d.gauge.Percent = int(data.UsageStat.UsedPercent)
	d.gauge.BorderLabel = fmt.Sprintf("disk usage (%dGB / %dGB) %d％",
		int(util.Byte2GBi(data.UsageStat.Used)),
		int(util.Byte2GBi(data.UsageStat.Total)),
		int(data.UsageStat.UsedPercent))
}

func (d *DiskUsageContainer) CreateUi() termui.GridBufferer {
	d.gauge = termui.NewGauge()
	d.gauge.Percent = 30
	d.gauge.Width = termui.TermWidth() / 2
	d.gauge.Height = 3
	d.gauge.BorderLabel = fmt.Sprint("disk usage")
	d.gauge.PercentColor = termui.ColorYellow
	d.gauge.BarColor = termui.ColorMagenta
	d.gauge.BorderFg = termui.ColorWhite
	d.gauge.BorderLabelFg = termui.ColorMagenta

	return d.gauge
}
