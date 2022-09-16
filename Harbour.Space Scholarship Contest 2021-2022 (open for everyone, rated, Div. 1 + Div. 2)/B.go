package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// couldn't solve
func (in *B) run() {
	t := in.nextInt()
	found := false
	for T := 0; T < t; T++ {
		s := in.nextString()
		t := in.nextString()
		for i := 0; i < len(s); i++ {
			if found {
				break
			}
			if s[i] != t[0] {
				continue
			} else {
				covered := 1
				went := 0
				for j := i + 1; j < len(t); j++ {
					if s[j] == t[covered] {
						covered++
						went = covered
						if covered == len(t) {
							found = true
							break
						}
					} else if j-2 >= 0 && s[j-2] == t[covered] {
						covered++
						went--
						j -= 2
						if went == 0 && covered == len(t) {
							found = true
							break
						}
					} else {
						break
					}
				}
			}
		}
		if found {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
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
	NewB(F).run()
}

func NewB(r *bufio.Reader) *B {
	return &B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
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
