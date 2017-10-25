package container

import (
	"fmt"
	dataservice "github.com/bunbunjp/gotop/dataservice/process"
	"github.com/gizak/termui"
	"math"
	"strings"
	"unicode/utf8"
	"log"
)

var rowHeaders = []string{"PID", "Name", "CPU(%)", "MEM(%)"}

// ProcessListContainer プロセス一覧を構成するコンテナーです
type ProcessListContainer struct {
	table       *termui.Table
	visibleRows *[][]string
}

// Initialize # Container Interface
func (p *ProcessListContainer) Initialize() {
}

// UpdateData # Container Interface
func (p *ProcessListContainer) UpdateData() {
}

func (p *ProcessListContainer) getDefaultRow() []string {
	return []string{"99999", "________________", "___", "___"}
}

func (p *ProcessListContainer) nameStrRounding(full string) string {
	limitLine := utf8.RuneCountInString(p.getDefaultRow()[1])
	nameRunes := strings.Split(full, "")

	if len(nameRunes) <= limitLine {
		return full
	}

	return strings.Join(nameRunes[:limitLine], "") + "..."
}

// UpdateRender # Container Interface
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

	ceil := int(math.Min(float64(visiblelimit+byas), float64(len(data.Processes))))

	log.Println("byas, ", byas)
	log.Println("visiblelimit+byas, ", visiblelimit+byas)
	log.Println("data.Processes, ", data.Processes)

	for _, process := range data.Processes[byas : ceil] {

		(*p.visibleRows)[count+1][0] = fmt.Sprint(process.Pid)
		(*p.visibleRows)[count+1][1] = fmt.Sprint(p.nameStrRounding(process.Name))
		(*p.visibleRows)[count+1][2] = fmt.Sprintf("%.1f", process.CPUPercent)
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
	for _, v := range rowHeaders {
		header = append(header, v)
	}
	header[int(data.GetSortKey())] += sortIcon
	(*p.visibleRows)[0] = header
}

func getHeight() int {
	return 17
}

// CreateUI # Container Interface
func (p *ProcessListContainer) CreateUI() termui.GridBufferer {

	p.visibleRows = &[][]string{rowHeaders}

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
