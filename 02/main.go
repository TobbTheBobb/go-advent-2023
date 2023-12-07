package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var red, green, blue = 12, 13, 14
var regamenr = regexp.MustCompile(`Game (\d+)`)

func main() {
	myInput_Structure := readFile()
	fmt.Printf("%#v\n", myInput_Structure)
	erg := solveTask1(red, green, blue, myInput_Structure)
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

func solveTask1(red int, green int, blue int, games []string) int {
	sum := 0
	for _, g := range games {
		valid := true
		//fmt.Println(g)
		gameID := regamenr.FindAllStringSubmatch(g, -1)[0][1]
		fmt.Println("Game " + gameID)
		gameNr,err := strconv.Atoi(gameID)
		check(err)
		gameRes := strings.TrimSpace(strings.Split(g,":")[1])
		fmt.Println("Result " + gameRes)
		gameRou := strings.Split(gameRes,";")
		for _,gr := range gameRou {
			gr = strings.TrimSpace(gr)
			//fmt.Print(" "+gr)
			draws := strings.Split(gr,",")
			for _,dc := range draws {
				dc = strings.TrimSpace(dc)
				//fmt.Print(" "+dc)
				c := strings.Split(dc," ")
				count,err := strconv.Atoi(c[0])
				check(err)
				colour := c[1]
				switch colour {
				case "red":
					if count > red {
						valid = false
					}
				case "green":
					if count > green {
						valid = false
					}
				case "blue":
					if count > blue {
						valid = false
					}
				}
			}
		}
		if valid {
			sum = sum+gameNr
		}
		fmt.Println(sum)
	}
	return sum
}