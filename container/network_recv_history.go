package container

import (
	"fmt"
	"github.com/bunbunjp/gotop/dataservice/network"
	"github.com/bunbunjp/gotop/util"
	"github.com/gizak/termui"
)

type NetworkRecvHistoryContainer struct {
	linecharts []termui.Sparkline
}

func (n *NetworkRecvHistoryContainer) Initialize() {
}

func (n *NetworkRecvHistoryContainer) UpdateRender() {
	data := network.GetInstace()

	n.linecharts[0].Data = data.RecvHistory
	n.linecharts[0].Title = fmt.Sprintf("Recv total %.2fMB (%.2fKB/s)",
		util.Byte2MBi(data.LatestStatus.BytesSent),
		util.Byte2KBi(data.SentPerSecond),
	)
}

func (n *NetworkRecvHistoryContainer) CreateUi() termui.GridBufferer {
	data := network.GetInstace()

	n.linecharts = make([]termui.Sparkline, 0)
	n.linecharts = append(n.linecharts, termui.NewSparkline())
	n.linecharts[0].Data = data.SentHistory
	n.linecharts[0].Height = 4
	n.linecharts[0].LineColor = termui.ColorCyan
	lines := termui.NewSparklines(n.linecharts...)
	lines.Height = 7
	return lines
}
