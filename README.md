# quickly load objects from glade with gotk3 ( gotk3 从glade快速加载对象)

## usage （使用）

### Define struct (定义 struct)
``` go
type MyApp struct {
    //If the widget name starts with Wd, it will be automatically loaded
    //如果widget名称以Wd开头，会自动加载
	WdWindow *gtk.Window 
	WdMsg    *gtk.Label
	WdBtn    *gtk.Button
}

//Must implement loader.Ptr interface{}
//必须实现 loader.Ptr interface{} 的Val方法
func (app *MyApp) Val() interface{} {
	return *app
}

//The method starting with Sig will automatically connect
//Sig开头的方法会自动connect
func (app *MyApp) SigBtnCk() {
	app.WdMsg.SetLabel("clicked")
}
```
### load and connect (加载和连接)

```go
	app := new(MyApp)
	build, _ := gtk.BuilderNewFromFile("ui.glade")
	loader.Map(app, build)     //load widgets
	loader.Connect(app, build) //connect signals
```
### example(例子)
./example