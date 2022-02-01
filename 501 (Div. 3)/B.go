package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (in *B) run() {
	n := in.nextInt()
	s := []byte(in.nextString())
	t := in.nextString()
	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		if s[i] != t[i] {
			p := -1
			for j := i + 1; j < n; j++ {
				if t[i] == s[j] {
					p = j
					break
				}
			}
			if p == -1 {
				fmt.Println(p)
				return
			}
			for j := p - 1; j >= i; j-- {
				s[j], s[j+1] = s[j+1], s[j]
				arr = append(arr, j+1)
			}
		}
	}
	if string(s) != t {
		fmt.Println(-1)
		return
	}
	fmt.Println(len(arr))
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		fmt.Print(" ")
	}
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
	file, err := os.Open("input.txt")
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
