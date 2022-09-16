package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (sc *B) run() {
	x := sc.nextInt()
	money := sc.nextInt()
	got := 0
	for i := 0; i < x; i++ {
		got += sc.nextInt()
	}
	if money == got {
		fmt.Println("Just enough")
	} else if money < got {
		fmt.Println("Success")
	} else {
		fmt.Println("More funding required")
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
	NewB(F).run()
}

func NewB(r *bufio.Reader) *B {
	return &B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

type B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (sc *B) nextLine() string {
	line, err := sc.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	sc.split = []string{}
	sc.index = 0
	return line
}

func (sc *B) load() {
	if len(sc.split) <= sc.index {
		sc.split = strings.Split(sc.nextLine(), sc.separator)
		sc.index = 0
	}
}

func (sc *B) nextInt() int {
	sc.load()
	val, _ := strconv.Atoi(strings.TrimSpace(sc.split[sc.index]))
	sc.index++
	return val
}

func (sc *B) nextInt64() int64 {
	sc.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(sc.split[sc.index]), 10, 64)
	sc.index++
	return val
}

func (sc *B) nextString() string {
	sc.load()
	val := strings.TrimSpace(sc.split[sc.index])
	sc.index++
	return val
}
