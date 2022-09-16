package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (in *A) run() {
	n := in.nextInt()
	k := in.nextInt()
	str := in.nextString()
	p := make([]int, n)
	for i := 1; i < n; i++ {
		j := p[i-1]
		for j > 0 && str[i] != str[j] {
			j = p[j-1]
		}
		if str[i] == str[j] {
			j++
		}
		p[i] = j
	}
	l := len(str) - p[n-1]
	for i := 0; i < k-1; i++ {
		fmt.Print(str[:l])
	}
	fmt.Print(str)

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

func (in *A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}

func (in *A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *A) nextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *A) nextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *A) nextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}
