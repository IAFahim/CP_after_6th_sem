package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverse(arr []byte, start, end int) {
	for i, j := start, end-1; i < end/2; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func (in *B) run() {
	n := in.nextInt()
	s := in.nextString()
	c := make([]byte, n)
	for i := 0; i < n; i++ {
		c[i] = s[i]
	}
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			reverse(c, 0, i)
		}
	}
	fmt.Println(string(c))
}

type B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}

func (in *B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *B) nextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *B) nextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *B) nextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

func NewB(r *bufio.Reader) *B {
	return &B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
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
