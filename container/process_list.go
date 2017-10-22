package container

import (
	"fmt"
	dataservice "github.com/bunbunjp/gotop/dataservice/process"
	"github.com/gizak/termui"
	"math"
	"strings"
	"unicode/utf8"
)

var ROW_HEADERS = []string{"PID", "Name", "CPU(%)", "MEM(%)"}

type ProcessListContainer struct {
	table       *termui.Table
	visibleRows *[][]string
}

func (p *ProcessListContainer) Initialize() {

}

func (pc *ProcessListContainer) UpdateData() {
}

func (pc *ProcessListContainer) getDefaultRow() []string {
	return []string{"99999", "________________", "___", "___"}
}

func (pc *ProcessListContainer) nameStrRounding(full string) string {
	limitLine := utf8.RuneCountInString(pc.getDefaultRow()[1])
	nameRunes := strings.Split(full, "")

	if len(nameRunes) <= limitLine {
		return full
	}

	return strings.Join(nameRunes[:limitLine], "") + "..."
}

func (p *ProcessListContainer) UpdateRender() {
	data := dataservice.GetInstance()
	visiblelimit := getHeight() - 3
	selectedIndex := data.GetSelectedIndex()
	byas := int(math.Max(0.0, float64((selectedIndex+1)-visiblelimit)))
	count := 0

	// 選択中の行をカラーリング
	for i := 0; i < visiblelimit; i++ {
		if i == (selectedIndex - byas) {
			p.table.BgColors[i+1] = termui.ColorGreen
			p.table.FgColors[i+1] = termui.ColorBlack
		} else {
			p.table.BgColors[i+1] = termui.ColorBlack
			p.table.FgColors[i+1] = termui.ColorWhite
		}
	}

	for _, process := range data.Processes[byas : visiblelimit+byas] {

		(*p.visibleRows)[count+1][0] = fmt.Sprint(process.Pid)
		(*p.visibleRows)[count+1][1] = fmt.Sprint(p.nameStrRounding(process.Name))
		(*p.visibleRows)[count+1][2] = fmt.Sprintf("%.1f", process.CpuPercent)
		(*p.visibleRows)[count+1][3] = fmt.Sprintf("%.1f", process.MemPercent)

		count++
	}

	for ; count > visiblelimit; count++ {
		(*p.visibleRows)[count+1] = p.getDefaultRow()
	}

	var sortIcon string
	if data.GetIsReverse() {
		sortIcon = " ▼"
	} else {
		sortIcon = " ▲"
	}

	header := []string{}
	for _, v := range ROW_HEADERS {
		header = append(header, v)
	}
	header[int(data.GetSortKey())] += sortIcon
	(*p.visibleRows)[0] = header
}

func getHeight() int {
	return 17
}

func (p *ProcessListContainer) CreateUi() termui.GridBufferer {

	p.visibleRows = &[][]string{ROW_HEADERS}

	for i := 0; i < getHeight()-3; i++ {
		row := p.getDefaultRow()
		*p.visibleRows = append(*p.visibleRows, row)
	}

	p.table = termui.NewTable()
	p.table.FgColor = termui.ColorWhite
	p.table.BgColor = termui.ColorDefault
	p.table.TextAlign = termui.AlignLeft
	p.table.Separator = false
	p.table.Rows = *p.visibleRows
	p.table.X = termui.TermWidth() / 2
	p.table.Width = termui.TermWidth() / 2
	p.table.Analysis()
	p.table.SetSize()
	p.table.Border = true

	return p.table
}
