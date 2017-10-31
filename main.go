package main

import (
	"github.com/bunbunjp/gotop/container"
	"github.com/bunbunjp/gotop/dataservice/cpu"
	"github.com/bunbunjp/gotop/dataservice/disk"
	"github.com/bunbunjp/gotop/dataservice/memory"
	"github.com/bunbunjp/gotop/dataservice/network"
	"github.com/bunbunjp/gotop/dataservice/process"
	"github.com/comail/colog"
	"github.com/jroimartin/gocui"
	"log"
	"os"
	"time"
)

// Container UIコンテナーインターフェイス
type Container interface {
	Initialize()
	UpdateRender(g *gocui.Gui) error
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

	file, _ := os.OpenFile("main.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
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

	for _, c := range containers {
		c.Initialize()

		g.SetLayout(c.CreateUI)
	}

	if err := g.SetKeybinding("process_list", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	go func() {
		for {
			for _, c := range containers {
				g.Execute(c.UpdateRender)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
