package main

import "fmt"

// GUIFactory 抽象工厂接口声明了一组能返回不同抽象产品的方法。这些产品属于同一个系列
// 且在高层主题或概念上具有相关性。同系列的产品通常能相互搭配使用。系列产
// 品可有多个变体，但不同变体的产品不能搭配使用。
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// WinFactory 具体工厂可生成属于同一变体的系列产品。工厂会确保其创建的产品能相互搭配
// 使用。具体工厂方法签名会返回一个抽象产品，但在方法内部则会对具体产品进行实例化。
type WinFactory struct{}

// MacFactory 每个具体工厂中都会包含一个相应的产品变体。
type MacFactory struct{}

func (f *WinFactory) CreateButton() Button {
	return &WinButton{}
}

func (f *WinFactory) CreateCheckbox() Checkbox {
	return &WinCheckbox{}
}

func (f *MacFactory) CreateButton() Button {
	return &MacButton{}
}

func (f *MacFactory) CreateCheckbox() Checkbox {
	return &MacCheckbox{}
}

// Button 系列产品中的特定产品必须有一个基础接口。所有产品变体都必须实现这个接口。
type Button interface {
	Paint()
}

// Checkbox 这是另一个产品的基础接口。所有产品都可以互动，但是只有相同具体变体的产
// 品之间才能够正确地进行交互。
type Checkbox interface {
	Paint()
}

// WinButton 具体产品由相应的具体工厂创建。
type WinButton struct{}
type MacButton struct{}
type WinCheckbox struct{}
type MacCheckbox struct{}

func (b *WinButton) Paint() {
	fmt.Println("根据 Windows 样式渲染按钮")
}

func (b *MacButton) Paint() {
	fmt.Println("根据 macOS 样式渲染按钮")
}

func (c *WinCheckbox) Paint() {
	fmt.Println("根据 Windows 样式渲染复选框")
}

func (c *MacCheckbox) Paint() {
	fmt.Println("根据 macOS 样式渲染复选框")
}

// Application 客户端代码仅通过抽象类型（GUIFactory、Button 和 Checkbox）使用工厂
// 和产品。这让你无需修改任何工厂或产品子类就能将其传递给客户端代码。
type Application struct {
	Fac    GUIFactory
	Butt   Button
	Checkb Checkbox
}

func (app *Application) CreateUI() {
	app.Butt = app.Fac.CreateButton()
	app.Checkb = app.Fac.CreateCheckbox()
}

func (app *Application) Paint() {
	app.Butt.Paint()
	app.Checkb.Paint()
}

func NewApplication(factory GUIFactory) Application {
	var app Application
	app.Fac = factory
	return app
}

func main() {
	confOS := "Windows"
	var factory GUIFactory
	if confOS == "Windows" {
		factory = new(WinFactory)
	} else {
		factory = new(MacFactory)
	}
	app := NewApplication(factory)
	app.CreateUI()
	app.Paint()
}
