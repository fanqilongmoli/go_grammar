package main

import (
	"time"
	"fmt"
)

func main() {
	//定时器是当你想要在未来某一个时刻执行一次使用的
	//打点器则是当你想要在固定的事件间隔重复执行准备的

	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
