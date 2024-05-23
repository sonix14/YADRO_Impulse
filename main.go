package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("enter data:")
	n, containers, err := ScanData()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if CanBeSorted(n, containers) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func CanBeSorted(n int, containers [][]int) bool {
	// calculating total number of balls of each color
	ballsPerColor := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ballsPerColor[j] += containers[i][j]
		}
	}

	// calculating total number of balls in each container
	ballsPerContainer := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ballsPerContainer[i] += containers[i][j]
		}
	}

	// sorting
	sort.Ints(ballsPerColor)
	sort.Ints(ballsPerContainer)

	// checking the correspondence of the total number of balls
	// of each color to the number that can be placed in containers
	for i := 0; i < n; i++ {
		if ballsPerColor[i] != ballsPerContainer[i] {
			return false
		}
	}
	return true
}

func ScanData() (int, [][]int, error) {
	var n int
	fmt.Scan(&n)
	if n < 1 || n > 100 {
		err := fmt.Errorf("wrong number of containers and colors! (%d)", n)
		return 0, nil, err
	}

	containers := make([][]int, n)
	for i := 0; i < n; i++ {
		containers[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&containers[i][j])
			if containers[i][j] < 0 || containers[i][j] > 1000000000 {
				err := fmt.Errorf("wrong number of balls of the same color in one container! (%v)", containers[i][j])
				return 0, nil, err
			}
		}
	}
	return n, containers, nil
}
