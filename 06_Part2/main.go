package main

import (
	"fmt"
)

var sumOfResults = 1

func main() {
	races := make([][]int,1)
	race1 := make([]int,2)
	race1[0] = 55999793
	race1[1] = 401148522741405
	races[0] = race1
	for i := range races {
		fmt.Printf("Race %v\n",i+1)
		sumOfResults = processRace(races[i][0],races[i][1])
	}
	fmt.Println(sumOfResults)
}


func processRace(a,b int) int {
	fmt.Println(a,b)
	wins := 0
	for i:=0;i<a;i++ {
		if i*(a-i) > b {
			//fmt.Printf("Wenn i=%v dann komme ich %vmm weit. - Wins: %v\n",i,i*(a-i),wins+1)
			wins = wins+1
		}
	}
	fmt.Printf("Anzahl Siege: %v\n",wins)
	return wins
}