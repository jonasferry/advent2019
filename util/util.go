package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadF(path string) (lines []string, err error) {
	// Read input file into variable d (data)
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("fopen: ", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		path, err := os.Getwd()
		if err != nil {
			fmt.Println("patherr: ", err)
		}
		fmt.Println("path: ", path)
		fmt.Println("fileerr: ", err)
	}
	return

	//Naive version
	//d, _ := ioutil.ReadFile("input.txt")
	//fmt.Print(string(d))	// test
}
