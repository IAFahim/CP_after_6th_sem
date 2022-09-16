package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

func (in *B) run(out *bufio.Writer) {

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
	NewB(F).run(bufio.NewWriter(os.Stdout))
}

func NewB(r *bufio.Reader) *B {
	return &B{
		in:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

type B struct {
	in        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *B) GetLine() string {
	line, err := in.in.ReadString('\n')
	if err != nil {
		Println("error line:", line, " err: ", err)
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
