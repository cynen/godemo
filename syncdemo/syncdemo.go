package syncdemo

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// 测试sync.Cond
func TestCond() {
	fmt.Println("Cond test...")

	c := sync.NewCond(&sync.Mutex{})
	ready := 0

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Second * time.Duration(rand.Int63n(10)))
			// 加锁更改等待条件
			c.L.Lock()
			ready++
			c.L.Unlock()

			fmt.Printf("运动员%d已准备就绪\n", i)
			// 广播唤醒等待者，这里可以使用Broadcast和Signal
			c.Signal()
		}(i)
	}
	// 当修改条件或者 wait() 时，必须加锁，保护 condition
	for ready != 10 {
		c.L.Lock()
		c.Wait()
		c.L.Unlock()
		fmt.Println("裁判员被唤醒一次")
	}

	fmt.Println("所有运动员都准备就绪，比赛开始。。。")

}

// TestMap 测试系统自带的Map,此Map并非线程安全的.
var testmap = make(map[string]int)

func TestMap() {
	wg := sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
func get(key string) int {
	return testmap[key]
}

func set(key string, value int) {
	testmap[key] = value
}

// SyncMap 测试 sync内置的Map. 此Map是线程安全的.
func SyncMap() {
	var m = sync.Map{}

	wg := sync.WaitGroup{}
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k = %v, v = %v \n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// 测试WaitGroupDemo
func WgTest1() {
	// 判断当前cpu是多少核的,设置最大开启线程数.
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 创建一个wg,可以添加任务,每完成一个任务,就标记 Done,
	// main函数主要作用就是判断是否完成了所有的任务.如果都完成了,就退出程序.
	wg := sync.WaitGroup{}
	wg.Add(0) // 设置初始任务数.

	for i := 0; i < 10; i++ {
		// 开启线程.
		go func(wg *sync.WaitGroup, index int) {
			// 每开启一个线程,就任务列表Add 一个.
			wg.Add(1)
			a := 1
			for i := 0; i < 10; i++ {
				a += i
			}

			fmt.Println(index, a)
			// 完成一个任务,就标记Done.
			// 实际执行的是: wg.Add(-1)
			wg.Done()
		}(&wg, i) // 函数的入参
	}
	// wg等待所有任务完成.否则阻塞.
	wg.Wait()
}

// sync包下的Lock测试
// 排它锁.
func Lock() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("lock:", i)
	}
}

// 读锁
func Rlock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println("rlock:", i)
	}
}

// 写锁.
func WLock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("wlock:", i)
	}
}

var mutex = new(sync.Mutex)

func StdOut(s string) {
	mutex.Lock()
	defer mutex.Unlock()
	for _, data := range s {
		fmt.Printf("%c", data)
	}
	fmt.Println()
}

func PersonDemo(s string) {
	StdOut(s)
}
