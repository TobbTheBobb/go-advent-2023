package main

import (
	"fmt"
	"os"
	"strings"
	"math"
)

func main() {
	myInput_Structure := readFile()
	fmt.Printf("%#v\n",myInput_Structure)
	processStructure(myInput_Structure)
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

func processStructure(structure []string) {
	sumOfResults := float64(0)
	for _,s := range structure {
		//fmt.Println(s)
		lists := strings.Split(s, "|")
		list1 := strings.TrimSpace(strings.Split(lists[0],":")[1])
		values2 := strings.Split(lists[1]," ")
		values1 := strings.Split(list1," ")
		m := make(map[string]bool)
		for _,value := range values1 {
			for _,value2 := range values2 {
				if !m[value] && value == value2 && value != ""{
					m[value] = true
				}
			}
		}
		fmt.Println(m)
		if len(m) >= 1 {
			if len(m) == 1{
				fmt.Println(float64(len(m)))
				sumOfResults = sumOfResults + float64(len(m))
			} else {
				fmt.Println(math.Pow(2,float64(len(m)-1)))
				sumOfResults = sumOfResults + math.Pow(2,float64(len(m)-1))
			}
			
			fmt.Println(sumOfResults)
		}
	}
	fmt.Printf("Ergebnis: %v\n", sumOfResults)
}