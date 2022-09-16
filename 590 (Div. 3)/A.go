package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func (sc *A) run() {
	t := sc.nextInt()
	for i := 0; i < t; i++ {
		n := sc.nextInt()
		sum := 0
		for i := 0; i < n; i++ {
			sum += sc.nextInt()
		}
		fmt.Println(int(math.Ceil(float64(sum) / float64(n))))
	}
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
	NewA(F).run()
}

func NewA(r *bufio.Reader) *A {
	return &A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

type A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (sc *A) nextLine() string {
	line, err := sc.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	sc.split = []string{}
	sc.index = 0
	return line
}

func (sc *A) load() {
	if len(sc.split) <= sc.index {
		sc.split = strings.Split(sc.nextLine(), sc.separator)
		sc.index = 0
	}
}

func (sc *A) nextInt() int {
	sc.load()
	val, _ := strconv.Atoi(strings.TrimSpace(sc.split[sc.index]))
	sc.index++
	return val
}

func (sc *A) nextInt64() int64 {
	sc.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(sc.split[sc.index]), 10, 64)
	sc.index++
	return val
}

func (sc *A) nextString() string {
	sc.load()
	val := strings.TrimSpace(sc.split[sc.index])
	sc.index++
	return val
}
