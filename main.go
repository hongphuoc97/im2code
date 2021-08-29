package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type Bar struct {
	startValue   float64
	currentValue int64
	finishValue  float64
	graph        string
	rate         float64
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	fmt.Println("Autobot: Nhap ten nguoi yeu cua ban")
	reader := bufio.NewReader(os.Stdin)
	girlFriendName, _ := reader.ReadString('\n')
	fmt.Println("You: " + girlFriendName)

	time.Sleep(time.Millisecond * 1000)

	fmt.Println("Autobot: really, That's not so cool")
	time.Sleep(time.Millisecond * 1000)

	fmt.Println("Autobot: okay, I'll find a lover for you")
	time.Sleep(time.Millisecond * 2000)
	CallClear()

	InitProgressBar(0, 100).play()
	CallClear()

	//ScanImageFile2Console(OpenFile("C:\\Users\\Phuoc\\Downloads\\ready.txt"), 30)
	time.Sleep(time.Millisecond * 500)
	CallClear()
	ScanImageFile2Console(OpenFile("C:\\Users\\Phuoc\\Downloads\\3.txt"), 30)
	time.Sleep(time.Millisecond * 500)
	CallClear()

	ScanImageFile2Console(OpenFile("C:\\Users\\Phuoc\\Downloads\\2.txt"), 30)
	time.Sleep(time.Millisecond * 500)
	CallClear()

	ScanImageFile2Console(OpenFile("C:\\Users\\Phuoc\\Downloads\\1.txt"), 30)
	time.Sleep(time.Millisecond * 500)
	CallClear()
	ScanImageFile2Console(OpenFile("C:\\Users\\Phuoc\\Downloads\\girl1.txt"), 30)
	CallClear()
	ScanImageFile2Console(OpenFile("C:\\Users\\Phuoc\\Downloads\\girl2.txt"), 30)
	for {
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func InitProgressBar(startValue, finishValue float64) *Bar {
	rate := float64(startValue / finishValue)
	bar := &Bar{
		startValue:  startValue,
		finishValue: finishValue,
		rate:        rate,
		graph:       "=",
	}

	return bar
}

func (bar *Bar) play() {
	bar.currentValue = int64(bar.rate * 100)
	for ; bar.currentValue <= 100; bar.currentValue++ {
		bar.printGraph()
		time.Sleep(time.Millisecond * 50)
	}
	//fmt.Println()
}

func (bar *Bar) printGraph() {
	space := strings.Repeat(" ", 100-int(bar.currentValue))
	graph := strings.Repeat(bar.graph, int(bar.currentValue)) + ">"
	searchIcon := "\\"
	if (bar.currentValue % 2) > 0 {
		searchIcon = "/"
	} else {
		searchIcon = "\\"
	}

	fmt.Printf("[%v%v] %v Searching %v%% \r", graph, space, searchIcon, bar.currentValue)
}

func OpenFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func ScanImageFile2Console(file *os.File, timeExcuteInMilisecond time.Duration) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		time.Sleep(time.Millisecond * timeExcuteInMilisecond)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}
