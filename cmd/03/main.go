package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var reDigit = regexp.MustCompile(`(?:\d+)+`)
var reNoDot = regexp.MustCompile(`[^.\d\n]`)

func main() {
	myInput_Structure := readFile()
	fmt.Printf("%#v\n", myInput_Structure)
	erg := solveTask1(myInput_Structure)
	fmt.Println(erg)
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
	myInput_Structure := strings.Split(input, "\n")
	return myInput_Structure
}


func doMatch(d []string) int {
	//fmt.Println("MATCH!")
	//fmt.Println(d[0])
	value,err := strconv.Atoi(d[0])
	check(err)
	return value
}

func solveTask1(lines []string) int {
	sum := 0
	for _, content := range lines {
		fmt.Println(content)
		digits := reDigit.FindAllStringSubmatch(content,-1)
		indexDigits := reDigit.FindAllStringSubmatchIndex(content,-1)
		indexNoDot := reNoDot.FindAllStringSubmatchIndex(content,-1)
		for c,d := range digits {
			matched := false
			for _,n := range indexNoDot {
				if ((n[0]==indexDigits[c][1] || n[1]==indexDigits[c][0]) && !matched){
					//fmt.Println(indexDigits[c])
					//fmt.Println(n)
					fmt.Println("Matched in same line. "+(digits[c][0]))
					sum = sum + doMatch(d)
					matched = true
				}
			}
		}
	}
	//fmt.Println(sum)
	//(\d+)[^.\d\n](\d+)?|(\d+)?[^.\d\n](\d+)
	return sum
}