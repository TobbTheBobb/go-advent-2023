package main

import (
	"strings"
	"os"
	//"regexp"
)

func main() {
	myInput_Structure := readFile()
	solveTask1(myInput_Structure)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile() []string {
	dat, err := os.ReadFile("./input")
	check(err)
	input := string(dat)
	//fmt.Println("Input-File:")
	//fmt.Println(input)
	myInput_Structure := strings.Split(input, "\n")
	return myInput_Structure
}

func solveTask1(structure []string) {

}