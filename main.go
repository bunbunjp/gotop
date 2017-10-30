package main

import (
	"github.com/bunbunjp/gotop/container"
	"github.com/bunbunjp/gotop/dataservice/cpu"
	"github.com/bunbunjp/gotop/dataservice/disk"
	"github.com/bunbunjp/gotop/dataservice/memory"
	"github.com/bunbunjp/gotop/dataservice/network"
	"github.com/bunbunjp/gotop/dataservice/process"
	"github.com/jroimartin/gocui"
	"log"
)

// Container UIコンテナーインターフェイス
type Container interface {
	Initialize()
	UpdateRender()
	CreateUI(g *gocui.Gui) error
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
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	containers := []Container{
		new(container.ProcessListContainer),
	}

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

	for _, c := range containers {
		c.Initialize()
		//err := c.CreateUI(g)

		g.SetLayout(c.CreateUI)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
