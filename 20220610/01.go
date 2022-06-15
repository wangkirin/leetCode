package main

import "fmt"

func main() {
	fmt.Println(GetSucessfulNum(3, [][]int{{0, 1, 1}, {0, 0, 2}, {1, 1, 1}, {1, 0, 2, 0, 0}, {3, 0, 2}}))
	fmt.Println(GetSucessfulNum(2, [][]int{{1, 17, 8, 11, 15}, {4, 10, 9}, {5, 10, 9, 14, 13}, {5, 14, 13, 10, 9}, {6, 14, 13}, {7, 14, 13}, {11, 14, 13}}))
}

func GetSucessfulNum(interval int, orders [][]int) int {
	count := 0
	seatOrdered := make(map[int]map[int]int)
	for _, order := range orders {
		time := order[0]/interval + 1
		if seatOrdered[time] == nil {
			seatOrdered[time] = make(map[int]int)
		}
		for i := 1; i < len(order); i += 2 {
			r, c := order[i], order[i+1]
			index := r*100 + c
			if _, ok := seatOrdered[time][index]; !ok {
				count++
				seatOrdered[time][index] = 1
				break
			}
		}
	}
	return count
}
