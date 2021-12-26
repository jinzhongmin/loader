package main

import (
	"loader"

	"github.com/gotk3/gotk3/gtk"
)

type MyApp struct {
	WdWindow *gtk.Window //Wd开头的gotk控件会自动加载
	WdMsg    *gtk.Label
	WdBtn    *gtk.Button
}

//Val  loader.Ptr interface{}需要实现的方法
func (app *MyApp) Val() interface{} {
	return *app
}

//Sig开头的方法会自动connect
func (app *MyApp) SigBtnCk() {
	app.WdMsg.SetLabel("clicked")
}

func main() {

	gtk.Init(nil)

	app := new(MyApp)
	build, _ := gtk.BuilderNewFromFile("ui.glade")
	loader.Map(app, build)     //load widgets
	loader.Connect(app, build) //connect signals

	app.WdWindow.SetTitle("example")
	app.WdWindow.Connect("destroy", func() {
		gtk.MainQuit()
	})

	app.WdWindow.ShowAll()

	gtk.Main()

}
