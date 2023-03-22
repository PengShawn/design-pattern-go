package main

import "fmt"

// Adaptee 已经存在的接口，需要被适配
type Adaptee struct{}

// SpecificRequest 示意方法，原本存在的方法
func (a *Adaptee) SpecificRequest() {
	// 具体的功能实现
	fmt.Println("具体的功能实现...")
}

// Target 定义客户端使用的接口，
type Target interface {
	Request()
}

// Adapter 对象适配器
type Adapter struct {
	adaptee Adaptee
}

func (a *Adapter) Request() {
	// 可能转调已经实现了的方法，进行适配
	a.adaptee.SpecificRequest()
}

func NewAdapter(adaptee Adaptee) Target {
	return &Adapter{
		adaptee: adaptee,
	}
}

// ClassAdapter 类适配器
type ClassAdapter struct {
	Adaptee
}

func (a *ClassAdapter) Request() {
	a.SpecificRequest()
}

func NewClassAdapter(adaptee Adaptee) Target {
	return &ClassAdapter{Adaptee: adaptee}
}

func main() {
	// 创建需要被适配的对象
	adaptee := Adaptee{}

	// 对象适配器
	target := NewAdapter(adaptee)
	target.Request()

	// 类适配器
	target1 := NewClassAdapter(adaptee)
	target1.Request()
}
