package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	fmt.Println("input row cols ticks respectively:")
	var M, N, ticks int
	_, err := fmt.Scan(&M, &N, &ticks)
	if err != nil {
		log.Fatal(err)
	}
	initialGrid := intialiseGrid(M, N)
	fmt.Println("intial Grid")
	printGrid(initialGrid)
	for i := 0; i < ticks; i++ {
		nextGrid := nextGeneration(initialGrid, M, N)
		fmt.Printf("After %d tick(s): \n", i+1)
		printGrid(nextGrid)
		initialGrid = nextGrid
	}

}

func intialiseGrid(M int, N int) [][]int {
	grid := makeArray(M, N)
	for k := 0; k < rand.Intn(M*N); k++ {
		i := rand.Intn(M)
		j := rand.Intn(N)
		grid[i][j] = 1
	}
	return grid
}

func makeArray(M int, N int) [][]int {
	grid := make([][]int, M)
	for i := range grid {
		grid[i] = make([]int, N)
	}
	return grid
}

func nextGeneration(intialGrid [][]int, M int, N int) [][]int {
	finalGrid := makeArray(M, N)
	for i := range intialGrid {
		for j := range intialGrid[i] {
			isLive := intialGrid[i][j] == 1
			liveN := countLiveNeighbors(intialGrid, i, j, M, N)
			if isLive && (liveN == 2 || liveN == 3) {
				finalGrid[i][j] = 1
			} else if !isLive && liveN == 3 {
				finalGrid[i][j] = 1
			} else {
				finalGrid[i][j] = 0
			}
		}
	}
	return finalGrid
}

func countLiveNeighbors(intialGrid [][]int, i int, j, M int, N int) int {
	countLive := 0
	if i < M-1 && intialGrid[i+1][j] == 1 {
		countLive++
	}
	if j < N-1 && intialGrid[i][j+1] == 1 {
		countLive++
	}
	if i < M-1 && j < N-1 && intialGrid[i+1][j+1] == 1 {
		countLive++
	}
	if i > 0 && intialGrid[i-1][j] == 1 {
		countLive++
	}
	if j > 0 && intialGrid[i][j-1] == 1 {
		countLive++
	}
	if i > 0 && j > 0 && intialGrid[i-1][j-1] == 1 {
		countLive++
	}
	return countLive
}

func printGrid(grid [][]int) {
	for i := range grid {
		fmt.Println(grid[i])
	}
}
