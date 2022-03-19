package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	val   int
	index int
}
type ByValIndex []Pair

func (e ByValIndex) Len() int {
	return len(e)
}

func (e ByValIndex) Less(i, j int) bool {
	if e[i].val == e[j].val {
		return e[i].index < e[j].index
	}
	return e[i].val < e[j].val
}

func (e ByValIndex) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (in *B) run() {
	n := in.nextInt()
	k := in.nextInt()

	pair := ByValIndex{}
	for i := 0; i < n; i++ {
		val := in.nextInt()
		pair = append(pair, Pair{
			val:   val,
			index: i,
		})
	}
	sort.Sort(pair)
	fmt.Printf("%v %d", pair, k)

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
