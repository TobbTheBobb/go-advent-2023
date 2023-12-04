package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var red, green, blue, power int
var regamenr = regexp.MustCompile(`Game (\d+)`)

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

func solveTask1(games []string) int {
	sum := 0
	for _, g := range games {
		//fmt.Println(g)
		gameID := regamenr.FindAllStringSubmatch(g, -1)[0][1]
		fmt.Println("Game " + gameID)
		gameRes := strings.TrimSpace(strings.Split(g,":")[1])
		fmt.Println("Result " + gameRes)
		gameRou := strings.Split(gameRes,";")
		red = 0
		green = 0
		blue = 0
		power = 0
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
						red = count
					}
				case "green":
					if count > green {
						green = count
					}
				case "blue":
					if count > blue {
						blue = count
					}
				}
			}
		}
		power = red * green * blue
		sum = sum+power
		fmt.Println(sum)
	}
	return sum
}