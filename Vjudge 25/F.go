package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	data string
	next map[string]string
}

func (n node) Search(key string) string {
	if n.data == key {
		return n.data
	}
	val, ok := n.next[key]
	if ok {
		return val
	}
	n.Search(key)
	return ""
}

func (sc *F) run() {
	n := sc.nextInt()
	for i := 0; i < n; i++ {
		cmd := sc.nextString()
		if cmd == "pwd" {

		} else {

		}
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
	NewF(F).run()
}

func NewF(r *bufio.Reader) *F {
	return &F{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

type F struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (sc *F) nextLine() string {
	line, err := sc.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	sc.split = []string{}
	sc.index = 0
	return line
}

func (sc *F) load() {
	if len(sc.split) <= sc.index {
		sc.split = strings.Split(sc.nextLine(), sc.separator)
		sc.index = 0
	}
}

func (sc *F) nextInt() int {
	sc.load()
	val, _ := strconv.Atoi(strings.TrimSpace(sc.split[sc.index]))
	sc.index++
	return val
}

func (sc *F) nextInt64() int64 {
	sc.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(sc.split[sc.index]), 10, 64)
	sc.index++
	return val
}

func (sc *F) nextString() string {
	sc.load()
	val := strings.TrimSpace(sc.split[sc.index])
	sc.index++
	return val
}
