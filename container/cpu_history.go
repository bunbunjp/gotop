package container

import (
	"github.com/jroimartin/gocui"
)

// CPUHistoryContainer CPU使用率を構成するコンテナー
type CPUHistoryContainer struct {
	//colors   []termui.Attribute
	//colorSet []termui.Attribute
	//lines    []termui.Sparkline
}

// Initialize # Container Interface
func (c *CPUHistoryContainer) Initialize() {
	//data := dataservice.GetInstance()
	//
	//c.colorSet = []termui.Attribute{termui.ColorCyan, termui.ColorMagenta, termui.ColorYellow, termui.ColorRed, termui.ColorGreen, termui.ColorBlue}
	//c.colors = make([]termui.Attribute, data.CoreCount, data.CoreCount)
	//
	//for i := 0; i < data.CoreCount; i++ {
	//	c.colors[i] = util.GetColorRand()
	//}
}

// UpdateRender # Container Interface
func (c *CPUHistoryContainer) UpdateRender() {
	//data := dataservice.GetInstance()
	//for i := 0; i < len(data.AccumuData); i++ {
	//	c.lines[i].Data = data.AccumuData[i]
	//	latestIdx := len(data.AccumuData[i]) - 1
	//	var percent int
	//	if latestIdx >= 0 {
	//		percent = data.AccumuData[i][latestIdx]
	//	} else {
	//		percent = 0
	//	}
	//	c.lines[i].Title = fmt.Sprintf("Core %d (%d％)", (i + 1), percent)
	//}

}

// CreateUI # Container Interface
func (c *CPUHistoryContainer) CreateUI(g *gocui.Gui) error {
	//data := dataservice.GetInstance()
	//
	//// single
	//c.lines = make([]termui.Sparkline, 0)
	//oneSparkLine := 2
	//for i := 0; i < len(data.AccumuData); i++ {
	//	spl := termui.NewSparkline()
	//	spl.Title = fmt.Sprintf("Core %d (%d％)", (i + 1), 0)
	//	spl.Height = oneSparkLine
	//	spl.LineColor = c.colorSet[i%len(c.colorSet)]
	//	c.lines = append(c.lines, spl)
	//}
	//
	//// group
	//group := termui.NewSparklines(c.lines...)
	//group.Height = len(data.AccumuData)*(oneSparkLine+1) + 2
	//group.Width = termui.TermWidth()
	//group.BorderLabel = "CPU availability"
	//
	//return group
	return nil
}
