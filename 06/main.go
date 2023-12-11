package main

import (
	"fmt"
)

var sumOfResults = 1

func main() {
	races := make([][]int,4)
	race1 := make([]int,2)
	race2 := make([]int,2)
	race3 := make([]int,2)
	race4 := make([]int,2)
	race1[0] = 55
	race1[1] = 401
	race2[0] = 99
	race2[1] = 1485
	race3[0] = 97
	race3[1] = 2274
	race4[0] = 93
	race4[1] = 1405
	races[0] = race1
	races[1] = race2
	races[2] = race3
	races[3] = race4
	for i := range races {
		fmt.Printf("Race %v\n",i+1)
		sumOfResults = sumOfResults * processRace(races[i][0],races[i][1])
	}
	fmt.Println(sumOfResults)
}


func processRace(a,b int) int {
	fmt.Println(a,b)
	wins := 0
	for i:=0;i<a;i++ {
		if i*(a-i) > b {
			fmt.Printf("Wenn i=%v dann komme ich %vmm weit. - Wins: %v\n",i,i*(a-i),wins+1)
			wins = wins+1
		}
	}
	fmt.Printf("Anzahl Siege: %v\n",wins)
	return wins
}