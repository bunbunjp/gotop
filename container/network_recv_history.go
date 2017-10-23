package container

import (
	"fmt"
	"github.com/bunbunjp/gotop/dataservice/network"
	"github.com/bunbunjp/gotop/util"
	"github.com/gizak/termui"
)

// NetworkRecvHistoryContainer ネットワーク受信通信率表示用コンテナ
type NetworkRecvHistoryContainer struct {
	linecharts []termui.Sparkline
}

// Initialize # Container Interface
func (n *NetworkRecvHistoryContainer) Initialize() {
}

// UpdateRender # Container Interface
func (n *NetworkRecvHistoryContainer) UpdateRender() {
	data := network.GetInstance()

	n.linecharts[0].Data = data.RecvHistory
	n.linecharts[0].Title = fmt.Sprintf("Recv total %.2fMB (%.2fKB/s)",
		util.Byte2MBi(data.LatestStatus.BytesSent),
		util.Byte2KBi(data.SentPerSecond),
	)
}

// CreateUI # Container Interface
func (n *NetworkRecvHistoryContainer) CreateUI() termui.GridBufferer {
	data := network.GetInstance()

	n.linecharts = make([]termui.Sparkline, 0)
	n.linecharts = append(n.linecharts, termui.NewSparkline())
	n.linecharts[0].Data = data.SentHistory
	n.linecharts[0].Height = 4
	n.linecharts[0].LineColor = termui.ColorCyan
	lines := termui.NewSparklines(n.linecharts...)
	lines.Height = 7
	return lines
}
