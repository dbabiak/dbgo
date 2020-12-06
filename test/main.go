package main

import (
	"fmt"
	"github.com/dbabiak/dbgo"
)

func main() {
	out, _, _ := dbgo.System2("ls -lrt ~ | grep Oct")
	for i, x := range dbgo.Lines(out) {
		fmt.Printf("%d %v\n", i, x)
	}
}
