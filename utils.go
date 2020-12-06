package dbgo

import (
	"bufio"
	"log"
	"os"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func PathToLines(pathname string) []string {
	f, err := os.Open(pathname)
	Check(err)
	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
