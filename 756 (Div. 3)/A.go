package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (in *A) run() {
	t := in.nextInt()
	arr := make([]bool, 128)
	arr['2'] = true
	arr['4'] = true
	arr['6'] = true
	arr['8'] = true
	for T := 0; T < t; T++ {
		str := in.nextString()
		if arr[str[len(str)-1]] {
			fmt.Println(0)
		} else if arr[str[0]] {
			fmt.Println(1)
		} else {
			x := -1
			for i := 0; i < len(str); i++ {
				if arr[str[i]] {
					x = 1
					if !arr[str[0]] && !arr[str[len(str)-1]] {
						x = 2
					}
				}
			}
			fmt.Println(x)
		}
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

//come on golang

func reverse(arr []byte, start, end int) {
	for i, j := start, end-1; i < end/2; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
