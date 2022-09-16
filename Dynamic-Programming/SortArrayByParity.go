package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sortArrayByParity(nums []int) []int {
	arr := make([][]int, 2)
	for i := 0; i < len(nums); i++ {
		arr[nums[i]%2] = append(arr[nums[i]%2], nums[i])
	}
	arr[0] = append(arr[0], arr[1]...)
	return arr[0]
}

func (sc *SortArrayByParity) run() {
	fmt.Printf("%v", sortArrayByParity([]int{3, 1, 2, 4}))
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
	NewSortArrayByParity(F).run()
}

func NewSortArrayByParity(r *bufio.Reader) *SortArrayByParity {
	return &SortArrayByParity{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

type SortArrayByParity struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (sc *SortArrayByParity) nextLine() string {
	line, err := sc.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	sc.split = []string{}
	sc.index = 0
	return line
}

func (sc *SortArrayByParity) load() {
	if len(sc.split) <= sc.index {
		sc.split = strings.Split(sc.nextLine(), sc.separator)
		sc.index = 0
	}
}

func (sc *SortArrayByParity) nextInt() int {
	sc.load()
	val, _ := strconv.Atoi(strings.TrimSpace(sc.split[sc.index]))
	sc.index++
	return val
}

func (sc *SortArrayByParity) nextInt64() int64 {
	sc.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(sc.split[sc.index]), 10, 64)
	sc.index++
	return val
}

func (sc *SortArrayByParity) nextString() string {
	sc.load()
	val := strings.TrimSpace(sc.split[sc.index])
	sc.index++
	return val
}
