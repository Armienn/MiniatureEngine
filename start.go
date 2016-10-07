package main

import "fmt"

func main() {
	fmt.Print("hahaee")
	b := make(map[string]int)
	fmt.Print(b["moskus"])
	b["moskus"] = 23
	fmt.Print(b["moskus"])
}
