package container

import (
	"fmt"
	"github.com/bunbunjp/gotop/dataservice/memory"
	"github.com/bunbunjp/gotop/util"
	"github.com/gizak/termui"
)

type SwapMemoryUsageContainer struct {
	swapGauge *termui.Gauge
}

func (s *SwapMemoryUsageContainer) Initialize() {
}

func (s *SwapMemoryUsageContainer) UpdateRender() {
	data := memory.GetInstance()

	s.swapGauge.Percent = int(data.LatestSwapStat.UsedPercent)
	s.swapGauge.BorderLabel = fmt.Sprintf("swap usage (%dMB/%dMB)",
		int(util.Byte2MBi(data.LatestSwapStat.Used)),
		int(util.Byte2MBi(data.LatestSwapStat.Total)))
}

func (s *SwapMemoryUsageContainer) CreateUI() termui.GridBufferer {

	s.swapGauge = termui.NewGauge()
	s.swapGauge.Width = termui.TermWidth() / 4
	s.swapGauge.Height = 10
	s.swapGauge.LabelAlign = termui.AlignRight

	return s.swapGauge
}
