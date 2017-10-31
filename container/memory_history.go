package container

import (
	"github.com/jroimartin/gocui"
)

// MemoryHistoryContainer メモリ使用率利用履歴表示用コンテナ
type MemoryHistoryContainer struct {
	//lineChart *termui.LineChart
}

// Initialize # Container Interface
func (m *MemoryHistoryContainer) Initialize() {

}

// UpdateRender # Container Interface
func (m *MemoryHistoryContainer) UpdateRender() {
	//data := memory.GetInstance()
	//
	//startline := int(math.Max(0, float64(len(data.VirtualUsedHistory)-m.lineChart.Width)))
	//m.lineChart.Data = data.VirtualUsedHistory[startline:]
}

// CreateUI # Container Interface
func (m *MemoryHistoryContainer) CreateUI(g *gocui.Gui) error {
	//m.lineChart = termui.NewLineChart()
	//m.lineChart.BorderLabel = fmt.Sprintf("memory usage history")
	//m.lineChart.Width = termui.TermWidth() / 2
	//m.lineChart.Height = 10
	//m.lineChart.AxesColor = termui.ColorWhite
	//m.lineChart.LineColor = termui.ColorGreen | termui.AttrBold
	//m.lineChart.DotStyle = '+'
	//m.lineChart.Mode = "dot"
	//
	//return m.lineChart

	return nil
}
