package concurrent

import (
	"fmt"
	"sync"
	"time"
)

// golang内存模型 http://ifeve.com/golang-mem/
// 官方GMM Blog bbbbbbbbbbnvbbbbhttps://golang.org/ref/mem

// http://bit.ly/2WLm8PU
// 理解互斥条件就行 mutual exclusion condition : 读锁共享,写锁只有一把,互斥 读写 写写 写读互斥
var (
	counter int            //计数器
	wg      sync.WaitGroup //信号量
)
var (
	rwMutex sync.RWMutex //读写锁
)

var (
	lock sync.Mutex            //互斥锁
	cond = sync.NewCond(&lock) //条件变量
)

func Mutex() {
	//1.两个信号
	wg.Add(1)

	//2.获取读锁
	fmt.Println("main thread wait rlock", time.Now())
	rwMutex.RLock()
	fmt.Println("main thread got rlock", time.Now())
	//3.开启子线程
	go incCounter()
	fmt.Println(counter)
	time.Sleep(time.Second * 5)
	//4.释放读锁
	rwMutex.RUnlock()
	fmt.Println("main thread release rlock", time.Now())

	//5.等待子线程结束
	wg.Wait()
	fmt.Println(counter)

}

func incCounter() {
	defer wg.Done()

	//2.1.获取写锁(独占锁)
	fmt.Println("sub thread wait rlock", time.Now())
	rwMutex.Lock()
	fmt.Println("sub thread got rlock", time.Now())

	//2.2.计数加1
	counter++

	//2.3.释放独占锁
	rwMutex.Unlock()
	fmt.Println("sub thread relese rlock", time.Now())

}

// 条件变量 http://bit.ly/2WHajdO
func Cond() {

	//1.获取锁
	cond.L.Lock()
	fmt.Println("main  thread got lock  ", time.Now())

	//2.开启线程执行一些事情
	go doSomething()

	//3.用与锁关联的条件变量的wait方法
	fmt.Println("main thread begin wait  ", time.Now())
	cond.Wait()
	fmt.Println("main thread end wait  ", time.Now())

	//4.释放锁
	cond.L.Unlock()
	fmt.Println("main thread release lock  ", time.Now())
}

func doSomething() {
	//2.1激活阻塞到条件变量的wait方法的一个线程

	time.Sleep(time.Second * 2)
	//2.2获取锁
	fmt.Println("sub thread begin get lock ", time.Now())
	//cond.L.Lock()
	fmt.Println("sub thread  got lock ", time.Now())

	time.Sleep(5 * time.Second)
	cond.Signal()

	//2.3释放锁
	//cond.L.Unlock()
	fmt.Println("sub thread release lock ", time.Now())
}

var a string
var once sync.Once

func setup() {
	time.Sleep(time.Second * 2) //1
	a = "hello, world"
	fmt.Println("setup over") //2
}

func doprint() {
	defer wg.Done()
	once.Do(setup) //3
	fmt.Println(a) //4
}

func twoprint() {
	go doprint()
	go doprint()
}

func Once() {
	wg.Add(2)
	twoprint()
	wg.Wait()
}

// once do用法和实现原理
// http://ifeve.com/go-concurrency-execute-once/
// http://bit.ly/2WPgm00  http://bit.ly/2WMKjxw
func OnceDo() {
	var num int
	sign := make(chan bool)
	var once sync.Once
	f := func(ii int) func() {
		return func() {
			num = (num + ii*2)
			sign <- true
		}
	}
	for i := 0; i < 3; i++ {
		fi := f(i + 1)
		go once.Do(fi)
	}
	for j := 0; j < 3; j++ {
		select {
		case <-sign:
			fmt.Println("Received a signal.")
		case <-time.After(100 * time.Millisecond):
			fmt.Println("Timeout!")
		}
	}
	fmt.Printf("Num: %d.\n", num)
}
