package main

func main() {

}

func canFinish(numCourses int, prerequisites [][]int) bool {
	depmap := make(map[int][]int)
	for _, prerequisite := range prerequisites {
		depmap[prerequisite[0]] = append(depmap[prerequisite[0]], prerequisite[1])
	}

}
