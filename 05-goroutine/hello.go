package main

import (
	"fmt"
	"sync"
	"time"
)

//@brief：耗时统计函数
func timeCost() func() {
	start := time.Now()
	return func() {
		tc:=time.Since(start)
		fmt.Printf("time cost = %d\n", tc)
	}
}

func foo(wg *sync.WaitGroup, sleep int) {
	defer timeCost()()		//注意，是对 timeCost()返回的函数进行调用，因此需要加两对小括号
    defer wg.Done()
    fmt.Println("This is function foo.")
	time.Sleep(time.Millisecond * time.Duration(sleep))  // 将 sleep 强制转换为 time.Duration 类型
}



func main() {
	defer timeCost()()		//注意，是对 timeCost()返回的函数进行调用，因此需要加两对小括号
    // 创建 WaitGroup 用于同步 Goroutine 的执行
    var wg sync.WaitGroup

    // 启动 3 个 Goroutine，每个 Goroutine 执行一个函数
    wg.Add(3)
    go foo(&wg,1)
    go foo(&wg,5)
    go foo(&wg,7)

    // 等待所有 Goroutine 执行完毕
    wg.Wait()
    fmt.Println("All functions have finished executing.")
}