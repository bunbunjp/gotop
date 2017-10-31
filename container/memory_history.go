package container

import (
	"fmt"
	"github.com/bunbunjp/gotop/dataservice/memory"
	"github.com/jroimartin/gocui"
	"log"
)

// MemoryHistoryContainer メモリ使用率利用履歴表示用コンテナ

const viewName = "memory_history"

type MemoryHistoryContainer struct {
	//lineChart *termui.LineChart
}

// Initialize # Container Interface
func (m *MemoryHistoryContainer) Initialize() {

}

// UpdateRender # Container Interface
func (m *MemoryHistoryContainer) UpdateRender(g *gocui.Gui) error {
	view, _ := g.View(viewName)
	log.Println()
	width, height := view.Size()

	data := memory.GetInstance()
	totalMem := data.TotalMemorySize

	oneScale := totalMem / height
	log.Println("oneScale is ", oneScale)
	log.Println(data.VirtualUsedHistory)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if y%2 == 0 {
				if x%2 == 0 {
					fmt.Fprint(view, "\033[41m\033[31m \033[0m")
				} else {
					fmt.Fprint(view, "\033[42m\033[32m \033[0m")
				}
			} else {
				if x%2 == 0 {
					fmt.Fprint(view, "\033[45m\033[33m \033[0m")
				} else {
					fmt.Fprint(view, "\033[46m\033[34m \033[0m")
				}
			}
		}
		fmt.Fprint(view, "\n")
	}
	//data := memory.GetInstance()
	//log.Println(data.VirtualUsedHistory)
	return nil
}

// CreateUI # Container Interface
func (m *MemoryHistoryContainer) CreateUI(g *gocui.Gui) error {
	//m.lineChart.Height = 10
	//m.lineChart.AxesColor = termui.ColorWhite
	//m.lineChart.LineColor = termui.ColorGreen | termui.AttrBold
	//m.lineChart.DotStyle = '+'
	//m.lineChart.Mode = "dot"
	//
	//return m.lineChart
	log.Println("MemoryHistoryContainer ... CreateUI")

	maxX, maxY := g.Size()
	if v, err := g.SetView(viewName, maxX/2, 0, maxX, (maxY/3)*1); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
			return err
		}
		v.Editable = false
		v.Wrap = true

		v.Title = viewName

		if err := g.SetCurrentView(viewName); err != nil {
			log.Panicln(err)
			return err
		}
	}

	return nil
}
