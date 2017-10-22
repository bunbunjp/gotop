package container


import (
	"github.com/gizak/termui"
	"github.com/bunbunjp/gotop/dataservice/memory"
	"fmt"
	"math"
)

type MemoryHistoryContainer struct {
	lineChart *termui.LineChart
}

func (m *MemoryHistoryContainer) Initialize() {

}

func (m *MemoryHistoryContainer) UpdateRender() {
	data := memory.GetInstance()

	startline := int(math.Max(0, float64(len(data.VirtualUsedHistory) - m.lineChart.Width)))
	m.lineChart.Data = data.VirtualUsedHistory[startline:]
}

func (m *MemoryHistoryContainer) CreateUi() termui.GridBufferer {
	m.lineChart = termui.NewLineChart()
	m.lineChart.BorderLabel = fmt.Sprintf("memory usage history")
	m.lineChart.Width = termui.TermWidth() / 2
	m.lineChart.Height = 10
	m.lineChart.AxesColor = termui.ColorWhite
	m.lineChart.LineColor = termui.ColorGreen | termui.AttrBold
	m.lineChart.DotStyle = '+'
	m.lineChart.Mode = "dot"

	return m.lineChart
}

