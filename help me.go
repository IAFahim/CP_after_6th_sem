package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err == nil {
		os.Stdin = file
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	b(os.Stdin)
}

//func a() {
//	n := 0
//	fmt.Scan(&n)
//	arr := make([]int, 7)
//	for j := 0; j < n; j++ {
//		for i := 0; i < 7; i++ {
//			fmt.Scan(&arr[i])
//
//		}
//		fmt.Println(arr[6]-arr[5], arr[5]-arr[1], arr[1])
//	}
//}

func b(file *os.File) {
	till := 0
	fmt.Scan(&till)
	reader := bufio.NewReader(file)
	for t := 0; t < till; t++ {
		n := 0
		fmt.Scan(&n)
		var c uint8
		str, _ := reader.ReadString('\n')
		str, _ = reader.ReadString('\n')
		str = strings.TrimSuffix(str, "\r\n")
		c = str[0]
		count := 1.0
		for i := 1; i < len(str); i++ {
			if str[i] == ' ' {
				continue
			}
			if c == str[i] {
				count++
			} else {
				x := int(math.Ceil(count / 2))
				for j := 0; j < x; j++ {
					print(string(c))
				}
				count = 0
				c = str[i]
			}
		}
		x := int(math.Ceil(count / 2))
		for j := 0; j < x; j++ {
			print(string(c))
		}
		print("\n")
	}
}
