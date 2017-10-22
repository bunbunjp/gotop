package container

import (
	"fmt"
	"github.com/bunbunjp/gotop/dataservice/memory"
	"github.com/gizak/termui"
	"math"
)

// MemoryHistoryContainer メモリ使用率利用履歴表示用コンテナ
type MemoryHistoryContainer struct {
	lineChart *termui.LineChart
}

// Initialize # Container Interface
func (m *MemoryHistoryContainer) Initialize() {

}

// UpdateRender # Container Interface
func (m *MemoryHistoryContainer) UpdateRender() {
	data := memory.GetInstance()

	startline := int(math.Max(0, float64(len(data.VirtualUsedHistory)-m.lineChart.Width)))
	m.lineChart.Data = data.VirtualUsedHistory[startline:]
}

// CreateUI # Container Interface
func (m *MemoryHistoryContainer) CreateUI() termui.GridBufferer {
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
