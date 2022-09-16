package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (in *A) run() {
	n, k := in.nextInt(), in.nextInt()
	type pair struct {
		v int
		i int
	}
	mp := map[int]pair{}
	for i := 0; i < n; i++ {
		x := in.nextInt()
		p := mp[x]
		p.v++
		p.i = i + 1
		mp[x] = p
	}
	if len(mp) >= k {
		fmt.Println("YES")
		for _, v := range mp {
			if k == 0 {
				break
			}
			fmt.Print(v.i)
			fmt.Print(" ")
			k--
		}
	} else {
		fmt.Println("NO")
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

func NewA(r *bufio.Reader) *A {
	return &A{
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
	NewA(F).run()
}
