package main

import (
	"github.com/bunbunjp/gotop/container"
	"github.com/bunbunjp/gotop/dataservice/cpu"
	"github.com/bunbunjp/gotop/dataservice/disk"
	"github.com/bunbunjp/gotop/dataservice/memory"
	"github.com/bunbunjp/gotop/dataservice/network"
	"github.com/bunbunjp/gotop/dataservice/process"
	"github.com/comail/colog"
	ui "github.com/gizak/termui"
	"log"
	"os"
	"time"
)

// Container UIコンテナーインターフェイス
type Container interface {
	Initialize()
	UpdateRender()
	CreateUI() ui.GridBufferer
}

// DataService システム情報のデータサービスインターフェイス
type DataService interface {
	Initialize()
}

type containerMap struct {
	m map[string]Container
}

func (m *containerMap) get(key string) Container {
	v, isOk := m.m[key]

	if isOk {
		return v
	}
	return nil
}

func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	file, err := os.OpenFile("main.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	colog.Register()
	colog.SetOutput(file)
	colog.ParseFields(true)
	colog.SetFormatter(&colog.StdFormatter{
		Flag: log.Lshortfile,
	})

	dataservices := []DataService{
		cpudata.GetInstance(),
		memory.GetInstance(),
		process.GetInstance(),
		network.GetInstance(),
		disk.GetInstance(),
	}

	for _, v := range dataservices {
		v.Initialize()
	}
	containers := containerMap{
		m: map[string]Container{
			"CpuHistory":         new(container.CPUHistoryContainer),
			"MemoryHistory":      new(container.MemoryHistoryContainer),
			"SwapMemory":         new(container.SwapMemoryUsageContainer),
			"VirtualMemory":      new(container.VirtualMemoryUsageContainer),
			"ProcessList":        new(container.ProcessListContainer),
			"NetworkSentHistory": new(container.NetworkSentHistoryContainer),
			"NetworkRecvHistory": new(container.NetworkRecvHistoryContainer),
			"DiskUsage":          new(container.DiskUsageContainer),
		},
	}

	for _, v := range containers.m {
		v.Initialize()
	}

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, containers.get("CpuHistory").CreateUI()),
		),
		ui.NewRow(
			ui.NewCol(6, 0, containers.get("MemoryHistory").CreateUI()),
			ui.NewCol(3, 0, containers.get("SwapMemory").CreateUI()),
			ui.NewCol(3, 0, containers.get("VirtualMemory").CreateUI()),
		),
		ui.NewRow(
			ui.NewCol(6, 0,
				containers.get("NetworkSentHistory").CreateUI(),
				containers.get("NetworkRecvHistory").CreateUI(),
				containers.get("DiskUsage").CreateUI(),
			),
			ui.NewCol(6, 0, containers.get("ProcessList").CreateUI()),
		),
	)

	//go updateGoroutine(containers)

	// handle key q pressing
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		// press q to quit
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/p", func(e ui.Event) {
		// handle all other key pressing
		process.GetInstance().ChangeSortKey(process.Pid)
		containers.get("ProcessList").UpdateRender()
	})

	ui.Handle("/sys/kbd/m", func(e ui.Event) {
		// handle all other key pressing
		process.GetInstance().ChangeSortKey(process.Memory)
		containers.get("ProcessList").UpdateRender()
	})

	ui.Handle("/sys/kbd/c", func(e ui.Event) {
		// handle all other key pressing
		process.GetInstance().ChangeSortKey(process.CPU)
		containers.get("ProcessList").UpdateRender()
	})

	ui.Handle("/sys/kbd/<down>", func(e ui.Event) {
		// handle all other key pressing
		process.GetInstance().IncrementSelectedIndex(+1)
		containers.get("ProcessList").UpdateRender()
	})

	ui.Handle("/sys/kbd/<up>", func(e ui.Event) {
		// handle all other key pressing
		process.GetInstance().IncrementSelectedIndex(-1)
		containers.get("ProcessList").UpdateRender()
	})

	// handle a 1s timer
	ui.Handle("/timer/1s", func(e ui.Event) {
		ui.Body.Align()
		for _, cont := range containers.m {
			//blocks = append(blocks, cont.CreateUi(&y))
			cont.UpdateRender()
		}
		ui.Render(ui.Body) // feel free to call Render, it's async and non-block
		time.Sleep(1 * time.Second)
	})

	ui.Loop() // block until StopLoop is called
	// event handler...
}
