package main

import (
	"fmt"
	"sync"
)

// Singleton 是单例模式接口，导出的
// 通过该接口可以避免 GetInstance 返回一个包私有类型的指针
type Singleton interface {
	foo()
}

type singleton struct{}

func (s singleton) foo() {}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func TestSingleton() {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		fmt.Println("两个实例不同")
	} else {
		fmt.Println("两个实例相同")
	}
}

const parCount = 100

func TestParallelSingleton() {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]Singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			// 协程阻塞，等待channel被关闭才能继续运行
			<-start
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	// 关闭channel，所有协程同时开始运行，实现并行(parallel)
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			fmt.Printf("协程测试：实例%v和%v不同\n", i, i-1)
			return
		}
	}
	fmt.Println("协程测试：所有实例相同")
}

func main() {
	TestSingleton()
	TestParallelSingleton()
}
