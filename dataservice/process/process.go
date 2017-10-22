package process

import (
	"time"
	"sort"
	"github.com/mattn/go-pipeline"
	"encoding/csv"
	"strings"
	"strconv"
	"path"
)

type MiniProcessData struct {
	Pid int32
	MemPercent float32
	CpuPercent float64
	Name string

}

/*********************************************/
/*********** Sort interfaces      ************/
/*********************************************/

type ProcessArray []MiniProcessData

func (p ProcessArray) Len() int {
	return len(p)
}

func (p ProcessArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type ByPID struct { ProcessArray }

func (bp ByPID) Less(i, j int) bool {
	return (bp.ProcessArray[i].Pid < bp.ProcessArray[j].Pid)
}

type ByMEM struct { ProcessArray }

func (bm ByMEM) Less(i, j int) bool {
	return (bm.ProcessArray[i].MemPercent < bm.ProcessArray[j].MemPercent)
}

type ByCPU struct { ProcessArray }

func (bc ByCPU) Less(i, j int) bool {
	return (bc.ProcessArray[i].CpuPercent < bc.ProcessArray[j].CpuPercent)
}

/*********************************************/

type SortKey int

const (
	Pid SortKey = iota
	Name
	Cpu
	Memory
)

type ProcessDataService struct {
	Processes ProcessArray
	isReverse bool
	sortKey SortKey
	isUpdateing bool
	selectedDataIndex int
}

var sharedInstance *ProcessDataService = &ProcessDataService{
	isReverse:true,
	sortKey:Cpu,
	isUpdateing:false,
}

func GetInstance() *ProcessDataService {
	return sharedInstance
}

func (p *ProcessDataService) IncrementSelectedIndex(incre int) {
	if (p.selectedDataIndex + incre <= 0) {
		return
	}
	p.selectedDataIndex += incre
}

func (p *ProcessDataService) GetSelectedIndex() int {
	return p.selectedDataIndex
}

func (p *ProcessDataService) GetSortKey() SortKey {
	return p.sortKey
}

func (p *ProcessDataService) GetIsReverse() bool {
	return p.isReverse
}

func (p *ProcessDataService) ChangeSortKey(key SortKey) {
	if key == p.sortKey {
		p.isReverse = !p.isReverse
		return
	}

	p.isReverse = false
	p.sortKey = key
}


func (p *ProcessDataService) Initialize() {
	go p.updateGoroutine()
}

func (p *ProcessDataService) updateGoroutine() {
	for {
		p.update()

		time.Sleep(2 * time.Second)
	}
}

func (p *ProcessDataService) sort(data ProcessArray) ProcessArray {
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
	case Cpu:
		fx(ByCPU{data})
		return data
	case Memory:
		fx(ByMEM{data})
		return data
	}
	return data
}

func (p *ProcessDataService) update() {
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
	tmp := ProcessArray{}
	for _, r := range records {
		pid, _ := strconv.ParseInt(r[0], 10, 32)
		cpu, _ := strconv.ParseFloat(r[1], 64)
		mem, _ := strconv.ParseFloat(r[2], 32)
		name := r[3]

		tmp = append(tmp, MiniProcessData{
			Pid		: int32(pid),
			CpuPercent	: cpu,
			MemPercent	: float32(mem),
			Name		: path.Base(name),
		})
 	}

	p.Processes = p.sort(tmp)
}
