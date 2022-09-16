package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	f, _ := ioutil.ReadFile("in.txt")
	str := strings.ReplaceAll(string(f), "\r\n", ",")
	print(str)

}
