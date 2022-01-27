package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

///xxxx
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func (in *C) run() {
	T := in.nextInt()
	for t := 0; t < T; t++ {
		a := in.nextString()
		s := in.nextString()
		i, j := len(a)-1, len(s)
		x := 1
		store := strings.Builder{}
		m := int(rune(s[j-1]) - '0')
		for i >= 0 && j > 0 {
			n := int(rune(a[i]) - '0')
			if i == 0 {
				m, _ = strconv.Atoi(s[0:j])
				store.WriteString(strconv.Itoa(m - n))
				i--
			} else if n > m {
				x++
				m = m*10 + int(s[j-x]-'0')
			} else {
				m, _ = strconv.Atoi(s[j-x : j])
				store.WriteString(strconv.Itoa(m - n))
				j -= x
				x = 1
				i--
			}
		}
		if i == -1 {
			num, _ := strconv.Atoi(reverse(store.String()))
			fmt.Println(num)
		} else {
			store.String()
			fmt.Println(-1)
		}
	}
}

type C struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *C) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}

func (in *C) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *C) nextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *C) nextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *C) nextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

func NewC(r *bufio.Reader) *C {
	return &C{
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
	NewC(F).run()
}
