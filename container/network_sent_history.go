package container

import (
	"fmt"
	"github.com/bunbunjp/gotop/dataservice/network"
	"github.com/bunbunjp/gotop/util"
	"github.com/gizak/termui"
)

// NetworkSentHistoryContainer ネットワーク送信通信料表示コンテナ
type NetworkSentHistoryContainer struct {
	linecharts []termui.Sparkline
}

// Initialize # Container Interface
func (n *NetworkSentHistoryContainer) Initialize() {
}

// UpdateRender # Container Interface
func (n *NetworkSentHistoryContainer) UpdateRender() {
	data := network.GetInstance()

	n.linecharts[0].Data = data.SentHistory
	n.linecharts[0].Title = fmt.Sprintf("Sent total %.2fMB (%.2fKB/s)",
		util.Byte2MBi(data.LatestStatus.BytesSent),
		util.Byte2KBi(data.SentPerSecond),
	)
	n.linecharts[0].TitleColor = termui.ColorCyan
}

// CreateUI # Container Interface
func (n *NetworkSentHistoryContainer) CreateUI() termui.GridBufferer {
	data := network.GetInstance()

	n.linecharts = make([]termui.Sparkline, 0)
	n.linecharts = append(n.linecharts, termui.NewSparkline())
	n.linecharts[0].Data = data.SentHistory
	n.linecharts[0].Height = 4
	n.linecharts[0].LineColor = termui.ColorBlue

	lines := termui.NewSparklines(n.linecharts...)
	lines.Height = 7
	lines.Width = termui.TermWidth() / 2
	return lines
}
