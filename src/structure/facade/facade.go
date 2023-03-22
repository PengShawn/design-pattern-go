package main

import "fmt"

type AModuleApi interface {
	TestA()
}

type AModuleImpl struct{}

func (a *AModuleImpl) TestA() {
	fmt.Println("现在在A模块里面操作TestA方法")
}

type BModuleApi interface {
	TestB()
}

type BModuleImpl struct{}

func (b *BModuleImpl) TestB() {
	fmt.Println("现在在B模块里面操作TestB方法")
}

type CModuleApi interface {
	TestC()
}

type CModuleImpl struct{}

func (c *CModuleImpl) TestC() {
	fmt.Println("现在在C模块里面操作TestC方法")
}

type Facade struct{}

func (f *Facade) Test() {
	a := AModuleImpl{}
	a.TestA()
	b := BModuleImpl{}
	b.TestB()
	c := CModuleImpl{}
	c.TestC()
}

func NewFacade() Facade {
	return Facade{}
}

func main() {
	facade := NewFacade()
	facade.Test()
}
