package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func (sc *E) run() {
	n := sc.nextInt()
	arr := make([]int, 5)
	for i := 0; i < n; i++ {
		arr[sc.nextInt()]++
	}
	taxi := arr[4]
	taxi += arr[3]
	arr[1] -= arr[3]
	taxi += arr[2] / 2
	if arr[2]%2 == 1 {
		taxi++
		arr[1] -= 2
	}
	if arr[1] > 0 {
		taxi += int(math.Ceil(float64(arr[1]) / 4))
	}
	fmt.Println(taxi)
}

func main() {
	file, err := os.Open("in.txt")
	var F *bufio.Reader
	if err == nil {
		F = bufio.NewReader(file)
		os.Stdin = file
	} else {
		F = bufio.NewReader(os.Stdin)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	NewE(F).run()
}

func NewE(r *bufio.Reader) *E {
	return &E{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

type E struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (sc *E) nextLine() string {
	line, err := sc.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	sc.split = []string{}
	sc.index = 0
	return line
}

func (sc *E) load() {
	if len(sc.split) <= sc.index {
		sc.split = strings.Split(sc.nextLine(), sc.separator)
		sc.index = 0
	}
}

func (sc *E) nextInt() int {
	sc.load()
	val, _ := strconv.Atoi(strings.TrimSpace(sc.split[sc.index]))
	sc.index++
	return val
}

func (sc *E) nextInt64() int64 {
	sc.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(sc.split[sc.index]), 10, 64)
	sc.index++
	return val
}

func (sc *E) nextString() string {
	sc.load()
	val := strings.TrimSpace(sc.split[sc.index])
	sc.index++
	return val
}
