package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex1 sync.Mutex
var mutex2 sync.Mutex

func routine1() {
	fmt.Println("Routine1: trying to lock mutex1")
	mutex1.Lock()
	fmt.Println("Routine1: locked mutex1")

	time.Sleep(1 * time.Second) // 模拟某些操作

	fmt.Println("Routine1: trying to lock mutex2")
	mutex2.Lock()
	fmt.Println("Routine1: locked mutex2")

	// 解锁
	mutex2.Unlock()
	mutex1.Unlock()
}

func routine2() {
	fmt.Println("Routine2: trying to lock mutex2")
	mutex2.Lock()
	fmt.Println("Routine2: locked mutex2")

	time.Sleep(1 * time.Second) // 模拟某些操作

	fmt.Println("Routine2: trying to lock mutex1")
	mutex1.Lock()
	fmt.Println("Routine2: locked mutex1")

	// 解锁
	mutex1.Unlock()
	mutex2.Unlock()
}

func main() {
	go routine1()
	go routine2()

	// 等待足够的时间以观察死锁现象
	time.Sleep(5 * time.Second)
	fmt.Println("Main: program finished (may be deadlocked)")
}
