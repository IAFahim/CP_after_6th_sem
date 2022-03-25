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
	for i := 0; i < t; i++ {
		n := in.nextInt()
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.nextInt()
		}
		flaw := false
		for i := 1; i < n; i++ {
			if !(arr[i] == arr[i-1]+1 || arr[i] == arr[i-1]-1 || (arr[i] == 1 && arr[i-1] == n) || (arr[i] == n && arr[i-1] == 1)) {
				flaw = true
				break
			}
		}
		if flaw {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")

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

func (in *A) nextLine() string {
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
		in.split = strings.Split(in.nextLine(), in.separator)
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
