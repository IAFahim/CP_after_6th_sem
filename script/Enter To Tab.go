package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	f, _ := ioutil.ReadFile("input.txt")
	str := strings.ReplaceAll(string(f), "\r\n", ",")
	print(str)

}
