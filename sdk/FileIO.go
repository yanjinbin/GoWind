package sdk

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sync"
)

func FileIO() {

	bytes, err := ioutil.ReadFile("./sdk/sample.txt")
	if err != nil {
		fmt.Println("reading file err, ", err)
		return
	}
	fmt.Println("打印文件全部内容")
	fmt.Println("内容:\n", string(bytes))
	fmt.Println("用flag再次打印")
	fptr := flag.String("fpath", "sample.txt", "file path to read from")
	flag.Parse()
	fmt.Println("file path is ", *fptr)
}

// 分块读取
// todo 命令行参数获取问题
func Bufio() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	fmt.Println("fpath路径", *fptr)
	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		_, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b))
	}
}

// 逐行读取
func BufScanner() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

//
func CreateFile() {
	fmt.Println("将字符串写入文件")
	file, err := os.Create("sample.json")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("close file err", err)
			return
		}
	}()
	n, err := file.WriteString("Hell Go")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(n, "byte write successful")
	fmt.Println("追加写文件")
	openFile, err := os.OpenFile("sample.json", os.O_APPEND|os.O_WRONLY, 0644)
	newLine := "File handling is easy."
	_, err = fmt.Fprintln(openFile, newLine)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consume(data chan int, done chan bool) {
	f, err := os.Create("concurrent")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

func ConcurrentWrite() {
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data, done)
	go func() {
		wg.Wait()
		close(data)
	}()
	d := <-done
	if d == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}
}
