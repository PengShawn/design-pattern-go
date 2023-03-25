package main

import "fmt"

type AbstractFactory interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}

type AbstractProductA interface {
	FuncOfProductA()
}

type AbstractProductB interface {
	FuncOfProductB()
}

type ProductA1 struct{}
type ProductA2 struct{}
type ProductB1 struct{}
type ProductB2 struct{}

func (p *ProductA1) FuncOfProductA() {
	fmt.Println("这是产品A1的方法")
}

func (p *ProductA2) FuncOfProductA() {
	fmt.Println("这是产品A2的方法")
}

func (p *ProductB1) FuncOfProductB() {
	fmt.Println("这是产品B1的方法")
}

func (p *ProductB2) FuncOfProductB() {
	fmt.Println("这是产品B2的方法")
}

type ConcreteFactory1 struct{}
type ConcreteFactory2 struct{}

func (f *ConcreteFactory1) CreateProductA() AbstractProductA {
	return &ProductA1{}
}

func (f *ConcreteFactory1) CreateProductB() AbstractProductB {
	return &ProductB1{}
}

func (f *ConcreteFactory2) CreateProductA() AbstractProductA {
	return &ProductA2{}
}

func (f *ConcreteFactory2) CreateProductB() AbstractProductB {
	return &ProductB2{}
}

func main() {
	var af AbstractFactory
	af = &ConcreteFactory1{}
	af.CreateProductA()
	af.CreateProductB()

	af = &ConcreteFactory2{}
	af.CreateProductA()
	af.CreateProductB()
}

