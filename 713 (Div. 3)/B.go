package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

func (in *B) run() {
	t := in.nextInt()
	for i := 0; i < t; i++ {
		n := in.nextInt()
		cm := make([][]byte, n)
		for i := 0; i < n; i++ {
			cm[i] = []byte(in.nextString())
		}
		point := []complex64{complex(0, 0), complex(0, 0)}
		c := 0
		for y := 0; y < n; y++ {
			for x := 0; x < n; x++ {
				if cm[y][x] == '*' {
					point[c] = complex(float32(x), float32(y))
					c++
				}
			}
		}

		if real(point[0]) == real(point[1]) {
			cm[(int(imag(point[0])))%n][(int(real(point[0]))+1)%n] = '*'
			cm[(int(imag(point[1])))%n][(int(real(point[1]))+1)%n] = '*'
		} else if imag(point[0]) == imag(point[1]) {
			cm[(int(imag(point[0]))+1)%n][(int(real(point[0])))%n] = '*'
			cm[(int(imag(point[1]))+1)%n][(int(real(point[1])))%n] = '*'
		} else {
			cm[(int(imag(point[0])))%n][(int(real(point[1])))] = '*'
			cm[(int(imag(point[1])))%n][(int(real(point[0])))] = '*'

		}
		for y := 0; y < n; y++ {
			Printf("%s\n", string(cm[y]))
		}
	}
}

func main() {
	var F *bufio.Reader
	in, inErr := os.Open("in.txt")
	if inErr == nil {
		F = bufio.NewReader(in)
	} else {
		F = bufio.NewReader(os.Stdin)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(in)
	NewB(F).run()
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
