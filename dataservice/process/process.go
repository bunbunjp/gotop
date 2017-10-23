package process

import (
	"encoding/csv"
	"github.com/mattn/go-pipeline"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
)

// MiniProcessData プロセスリストのソート用構造体
type MiniProcessData struct {
	Pid        int32
	MemPercent float32
	CPUPercent float64
	Name       string
}

/*********************************************/
/*********** Sort interfaces      ************/
/*********************************************/

// Processes プロセスリストの配列構造体
type Processes []MiniProcessData

// Len with Sort interface
func (p Processes) Len() int {
	return len(p)
}

// Swap with Sort interface
func (p Processes) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// ByPID プロセスリストのソートをPIDで行う
type ByPID struct{ Processes }

// Less with Sort interface
func (bp ByPID) Less(i, j int) bool {
	return (bp.Processes[i].Pid < bp.Processes[j].Pid)
}

// ByMEM プロセスリストのソートをメモリ利用率で行う
type ByMEM struct{ Processes }

// Less with Sort interface
func (bm ByMEM) Less(i, j int) bool {
	return (bm.Processes[i].MemPercent < bm.Processes[j].MemPercent)
}

// ByCPU プロセスリストのソートをCPU利用率で行う
type ByCPU struct{ Processes }

// Less with Sort interface
func (bc ByCPU) Less(i, j int) bool {
	return (bc.Processes[i].CPUPercent < bc.Processes[j].CPUPercent)
}

/*********************************************/

// SortKey 配列のキー識別子
type SortKey int

const (
	// Pid プロセスID
	Pid SortKey = iota

	// Name プロセス名
	Name

	// CPU CPU利用率
	CPU

	// Memory メモリ利用率
	Memory
)

// Service プロセスリストのデータサービス
type Service struct {
	Processes         Processes
	isReverse         bool
	sortKey           SortKey
	isUpdateing       bool
	selectedDataIndex int
}

var sharedInstance = &Service{
	isReverse:   true,
	sortKey:     CPU,
	isUpdateing: false,
}

// GetInstance is get singleton instance
func GetInstance() *Service {
	return sharedInstance
}

// IncrementSelectedIndex 選択中カラムの位置をインクリメントする
func (p *Service) IncrementSelectedIndex(incre int) {
	if p.selectedDataIndex+incre <= 0 {
		return
	}
	p.selectedDataIndex += incre
}

// GetSelectedIndex 選択中カラムの位置を取得する
func (p *Service) GetSelectedIndex() int {
	return p.selectedDataIndex
}

// GetSortKey 現在のソートキーを取得する
func (p *Service) GetSortKey() SortKey {
	return p.sortKey
}

// GetIsReverse 現在の並び順を取得する
func (p *Service) GetIsReverse() bool {
	return p.isReverse
}

// ChangeSortKey ソートキーを変更する
func (p *Service) ChangeSortKey(key SortKey) {
	if key == p.sortKey {
		p.isReverse = !p.isReverse
		return
	}

	p.isReverse = false
	p.sortKey = key
}

// Initialize with DataService interface
func (p *Service) Initialize() {
	go p.updateGoroutine()
}

func (p *Service) updateGoroutine() {
	for {
		p.update()

		time.Sleep(2 * time.Second)
	}
}

func (p *Service) sort(data Processes) Processes {
	fx := func(inData sort.Interface) {
		if p.isReverse {
			sort.Sort(sort.Reverse(inData))
		} else {
			sort.Sort(inData)
		}
	}

	switch p.sortKey {
	case Pid:
		fx(ByPID{data})
		return data
	case CPU:
		fx(ByCPU{data})
		return data
	case Memory:
		fx(ByMEM{data})
		return data
	}
	return data
}

func (p *Service) update() {
	if p.isUpdateing {
		return
	}
	p.isUpdateing = true
	defer func() {
		p.isUpdateing = false
	}()

	out, _ := pipeline.Output(
		[]string{"ps", "auxww"},
		[]string{"awk", "{print $2, $3, $4, $11}"},
	)

	strout := string(out)

	reader := csv.NewReader(strings.NewReader(strout))
	reader.Comma = ' '

	records, _ := reader.ReadAll()

	p.Processes = []MiniProcessData{}
	tmp := Processes{}
	for _, r := range records {
		pid, _ := strconv.ParseInt(r[0], 10, 32)
		cpu, _ := strconv.ParseFloat(r[1], 64)
		mem, _ := strconv.ParseFloat(r[2], 32)
		name := r[3]

		tmp = append(tmp, MiniProcessData{
			Pid:        int32(pid),
			CPUPercent: cpu,
			MemPercent: float32(mem),
			Name:       path.Base(name),
		})
	}

	p.Processes = p.sort(tmp)
}
