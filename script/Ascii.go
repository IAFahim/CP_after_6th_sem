package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (in *Ascii) run() {
	f, err := os.OpenFile("E:\\CSE332\\Assigment\\random input.txt", os.O_TRUNC|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("couldn't open")
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	fmt.Print(f.Name())
	f.WriteString("v2.0 raw\n")
	x := in.nextLine()
	for i := 0; i < len(x); i++ {
		str := strconv.FormatInt(int64(x[i]), 16) + " "
		f.WriteString(str)
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
	NewAscii(F).run()
}

func NewAscii(r *bufio.Reader) *Ascii {
	return &Ascii{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

type Ascii struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *Ascii) nextLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}

func (in *Ascii) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.nextLine(), in.separator)
		in.index = 0
	}
}

func (in *Ascii) nextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *Ascii) nextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *Ascii) nextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}
