package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
	"unsafe"
)

// 全局变量命名不规范
var global_var int
var GlobalVar2 string

// 常量命名错误
const max_SIZE = 100
const minsize = 10

// 1. 内存泄漏 - goroutine泄漏
func memoryLeak() {
	ch := make(chan int)
	go func() {
		for {
			select {
			case <-ch:
				return
			default:
				// 这个goroutine永远不会退出
				time.Sleep(time.Second)
			}
		}
	}()
	// 忘记关闭channel
}

// 2. 竞态条件
var counter int

func raceCondition() {
	for i := 0; i < 1000; i++ {
		go func() {
			counter++ // 没有锁保护
		}()
	}
}

// 3. 错误的错误处理
func badErrorHandling() {
	file, err := os.Open("nonexistent.txt")
	if err != nil {
		// 错误：忽略错误
		return
	}
	defer file.Close()

	data := make([]byte, 100)
	n, err := file.Read(data)
	// 错误：没有检查错误
	fmt.Println("Read", n, "bytes")
}

// 4. 资源泄漏
func resourceLeak() {
	file, _ := os.Open("test.txt") // 忽略错误
	// 错误：没有defer close
	data := make([]byte, 100)
	file.Read(data)
	// file永远不会被关闭
}

// 5. 切片越界和nil指针
func sliceAndNilErrors() {
	var slice []int
	fmt.Println(slice[0]) // panic: 空切片访问

	var ptr *int
	fmt.Println(*ptr) // panic: nil指针解引用

	arr := []int{1, 2, 3}
	fmt.Println(arr[10]) // panic: 索引越界
}

// 6. map并发读写
var sharedMap = make(map[string]int)

func mapConcurrencyIssue() {
	go func() {
		for {
			sharedMap["key"] = 1 // 并发写
		}
	}()

	go func() {
		for {
			_ = sharedMap["key"] // 并发读
		}
	}()
}

// 7. 无限循环
func infiniteLoop() {
	for {
		// 没有退出条件
		fmt.Println("Running forever...")
	}
}

// 8. 不当的类型转换
func badTypeConversion() {
	var i int64 = 1000000000000
	var j int32 = int32(i) // 数据截断
	fmt.Println(j)

	str := "not a number"
	num, _ := strconv.Atoi(str) // 忽略转换错误
	fmt.Println(num)
}

// 9. 字符串操作性能问题
func inefficientStringOps() {
	var result string
	for i := 0; i < 10000; i++ {
		result += "a" // 每次都创建新字符串
	}
	return
}

// 10. unsafe包的不当使用
func unsafeUsage() {
	s := "hello"
	ptr := unsafe.Pointer(&s)
	i := (*int)(ptr) // 类型不匹配的转换
	fmt.Println(*i)
}

// 11. 函数命名和设计问题
func do_something_bad(x int, y int, z int, a int, b int, c int) int { // 参数过多，命名不规范
	if x > 0 {
		if y > 0 {
			if z > 0 {
				if a > 0 {
					if b > 0 {
						if c > 0 {
							return x + y + z + a + b + c // 嵌套过深
						}
					}
				}
			}
		}
	}
	return 0
}

// 12. 接口设计问题
type BadInterface interface {
	DoEverything(string, int, bool, []byte, map[string]interface{}) (interface{}, error) // 过于复杂
	Method1()
	Method2()
	Method3()
	Method4()
	Method5() // 方法太多
}

// 13. 结构体设计问题
type BadStruct struct {
	data []byte // 未导出字段但应该是公开的
	Id   int    // 大小写混乱
	NAME string
	age  int
	// 没有适当的getter/setter
}

// 14. 死锁示例
func deadlockExample() {
	mu1 := &sync.Mutex{}
	mu2 := &sync.Mutex{}

	go func() {
		mu1.Lock()
		time.Sleep(time.Second)
		mu2.Lock() // 可能导致死锁
		mu2.Unlock()
		mu1.Unlock()
	}()

	go func() {
		mu2.Lock()
		time.Sleep(time.Second)
		mu1.Lock() // 可能导致死锁
		mu1.Unlock()
		mu2.Unlock()
	}()
}

// 15. 不正确的defer使用
func badDeferUsage() {
	for i := 0; i < 10; i++ {
		file, _ := os.Open(fmt.Sprintf("file%d.txt", i))
		defer file.Close() // 错误：在循环中defer
	}
}

// 16. 错误的接口实现
type Writer interface {
	Write([]byte) error
}

type BadWriter struct{}

func (w BadWriter) Write(data []byte) error {
	// 错误：没有实际写入数据，但返回nil
	return nil
}

// 17. 不合理的panic使用
func badPanicUsage(input string) {
	if input == "" {
		panic("input cannot be empty") // 应该返回错误而不是panic
	}
	fmt.Println(input)
}

// 18. 内存对齐问题
type BadAlignment struct {
	a bool  // 1 byte
	b int64 // 8 bytes
	c bool  // 1 byte
	d int64 // 8 bytes
	// 内存浪费严重
}

// 19. 错误的上下文使用
func badContextUsage() {
	// 应该使用context.Context但没有使用
	time.Sleep(10 * time.Second) // 阻塞操作没有超时控制
}

// 20. 不当的全局状态
var globalState = make(map[string]interface{}) // 全局可变状态

func modifyGlobalState(key string, value interface{}) {
	globalState[key] = value // 没有锁保护
}

// main函数也有问题
func Tttt() {
	// 错误：没有错误处理
	memoryLeak()
	raceCondition()
	badErrorHandling()

	// 错误：调用可能panic的函数但没有recover
	sliceAndNilErrors()

	// 错误：启动了并发操作但没有等待
	mapConcurrencyIssue()

	fmt.Println("Program completed") // 实际上可能已经panic了
}
