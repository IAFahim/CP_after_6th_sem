package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (sc *G) run() {
	//n, m, k := sc.nextInt(), sc.nextInt(), sc.nextInt()
	//count := 0
	//a := int(math.Min(float64(n), float64(k)))
	//b := int(math.Min(float64(m/2), float64(k)))
	//
	//fmt.Println(count)
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
	NewG(F).run()
}

func NewG(r *bufio.Reader) *G {
	return &G{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

type G struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (sc *G) nextLine() string {
	line, err := sc.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	sc.split = []string{}
	sc.index = 0
	return line
}

func (sc *G) load() {
	if len(sc.split) <= sc.index {
		sc.split = strings.Split(sc.nextLine(), sc.separator)
		sc.index = 0
	}
}

func (sc *G) nextInt() int {
	sc.load()
	val, _ := strconv.Atoi(strings.TrimSpace(sc.split[sc.index]))
	sc.index++
	return val
}

func (sc *G) nextInt64() int64 {
	sc.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(sc.split[sc.index]), 10, 64)
	sc.index++
	return val
}

func (sc *G) nextString() string {
	sc.load()
	val := strings.TrimSpace(sc.split[sc.index])
	sc.index++
	return val
}
