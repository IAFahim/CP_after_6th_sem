package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

type arrWithIndex struct {
	count int
	idx   int
}

func (in *A) run(out *bufio.Writer) {
	T := in.nextInt()
	for t := 0; t < T; t++ {
		n := in.nextInt()
		mp := map[int]arrWithIndex{}
		for i := 0; i < n; i++ {
			x := in.nextInt()
			mp[x] = arrWithIndex{count: mp[x].count + 1, idx: i}
		}

		for _, v := range mp {
			if v.count == 1 {
				Fprintln(out, v.idx+1)
				Println(v.idx + 1)
				break
			}
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
	NewA(F).run(bufio.NewWriter(os.Stdout))
}

func NewA(r *bufio.Reader) *A {
	return &A{
		in:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

type A struct {
	in        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *A) GetLine() string {
	line, err := in.in.ReadString('\n')
	if err != nil {
		Println("error line:", line, " err: ", err)
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
