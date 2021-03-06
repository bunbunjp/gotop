package container

import (
	"fmt"
	"github.com/bunbunjp/gotop/dataservice/memory"
	"github.com/bunbunjp/gotop/util"
	"github.com/gizak/termui"
)

// SwapMemoryUsageContainer メモリのスワップ領域の表示用コンテナ
type SwapMemoryUsageContainer struct {
	swapGauge *termui.Gauge
}

// Initialize # Container Interface
func (s *SwapMemoryUsageContainer) Initialize() {
}

// UpdateRender # Container Interface
func (s *SwapMemoryUsageContainer) UpdateRender() {
	data := memory.GetInstance()

	s.swapGauge.Percent = int(data.LatestSwapStat.UsedPercent)
	s.swapGauge.BorderLabel = fmt.Sprintf("swap usage (%dMB/%dMB)",
		int(util.Byte2MBi(data.LatestSwapStat.Used)),
		int(util.Byte2MBi(data.LatestSwapStat.Total)))
}

// CreateUI # Container Interface
func (s *SwapMemoryUsageContainer) CreateUI() termui.GridBufferer {

	s.swapGauge = termui.NewGauge()
	s.swapGauge.Width = termui.TermWidth() / 4
	s.swapGauge.Height = 10
	s.swapGauge.LabelAlign = termui.AlignRight

	return s.swapGauge
}
