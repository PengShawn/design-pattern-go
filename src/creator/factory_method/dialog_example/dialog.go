package main

import (
	"fmt"
	"runtime"
)

// Button 产品接口中将声明所有具体产品都必须实现的操作
type Button interface {
	Render()
	OnClick(f string)
}

// WindowsButton 具体产品需提供产品接口的各种实现
type WindowsButton struct {
}

func (b *WindowsButton) Render() {
	// 根据 Windows 样式渲染按钮
	fmt.Println("根据 Windows 样式渲染按钮")
}

func (b *WindowsButton) OnClick(f string) {
	// 绑定本地操作系统点击事件
	fmt.Printf("Windows按钮点击事件, f = %v\n", f)
}

type HTMLButton struct {
}

func (b *HTMLButton) Render() {
	// 根据 HTML 样式渲染按钮
	fmt.Println("根据 HTML 样式渲染按钮")
}

func (b *HTMLButton) OnClick(f string) {
	// 绑定网络浏览器的点击事件
	fmt.Printf("HTML按钮点击事件, f = %v\n", f)
}

// Dialog 创建者类声明的工厂方法必须返回一个产品类的对象。创建者的子类通常会提供该方法的实现。
type Dialog interface {
	CreateButton() Button
	Render()
}

type WindowsDialog struct{}

func (d *WindowsDialog) CreateButton() Button {
	return &WindowsButton{}
}

func (d *WindowsDialog) Render() {
	okButton := d.CreateButton()
	okButton.OnClick("closeDialog")
	okButton.Render()
}

type WebDialog struct{}

func (d *WebDialog) CreateButton() Button {
	return &HTMLButton{}
}

func (d *WebDialog) Render() {
	okButton := d.CreateButton()
	okButton.OnClick("closeDialog")
	okButton.Render()
}

func main() {
	var dialog Dialog
	// 程序根据当前配置或环境设定选择创建者的类型
	sysType := runtime.GOOS
	if sysType == "windows" {
		dialog = new(WindowsDialog)
	} else {
		dialog = new(WebDialog)
	}
	dialog.Render()

}
