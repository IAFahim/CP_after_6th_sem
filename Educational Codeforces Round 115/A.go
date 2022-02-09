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
	for T := 0; T < t; T++ {
		n := in.nextInt()
		arr := make([][]byte, 2, n)
		for y := 0; y < len(arr); y++ {
			arr[y] = []byte(in.nextString())
		}
		dfs := newDfs(arr, '0', '1')
		dfs.possible(0, 0)
		if dfs.found {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

type Dfs struct {
	dX       []int
	dY       []int
	grid     [][]byte
	find     byte
	obstacle byte
	found    bool
}

func newDfs(grid [][]byte, find byte, obstacle byte) *Dfs {
	dX := []int{1, 1, 1, 0, 0, -1, -1, -1}
	dY := []int{1, 0, -1, 1, -1, 1, 0, -1}
	return &Dfs{dX: dX, dY: dY, grid: grid, find: find, obstacle: obstacle, found: false}
}

func (d *Dfs) possible(x, y int) {
	for i := 0; i < len(d.dX); i++ {
		_y := y + d.dY[i]
		_x := x + d.dX[i]
		if 0 <= _y && _y < len(d.grid) && 0 <= _x && _x < len(d.grid[_y]) && d.grid[_y][_x] == d.find {
			if d.reached(_x, _y) {
				return
			}
			d.grid[_y][_x] = d.obstacle
			d.possible(_x, _y)
		}
	}
}

func (d *Dfs) reached(_x, _y int) bool {
	return _y == len(d.grid)-1 && _x == len(d.grid[_y])-1
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
