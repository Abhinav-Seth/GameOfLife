package main

import (
	model "GameOfLife/model"
	"GameOfLife/service"
	"fmt"
	"log"
)

func main() {
	fmt.Println("input row cols ticks respectively:")
	var M, N, ticks int
	_, err := fmt.Scan(&M, &N, &ticks)
	if err != nil {
		log.Fatal(err)
	}
	//var pts []model.Position
	//block
	//M = 4
	//N = 4
	//pts = []model.Position{{1, 1}, {1, 2}, {2, 1}, {2, 2}}

	//blinker
	//M = 5
	//N = 5
	//pts = []model.Position{
	//	{2, 1},
	//	{2, 2},
	//	{2, 3},
	//}
	//tub
	//M = 5
	//N = 5
	//pts = []model.Position{{1, 2}, {2, 1}, {2, 3}, {3, 2}}
	//Beacon
	//M = 6
	//N = 6
	//pts = []model.Position{{1, 1}, {1, 2}, {2, 1},
	//	{2, 2}, {3, 3}, {3, 4}, {4, 3}, {4, 4}}
	//M = 8
	//N = 8
	//pts = []model.Position{{2, 3}, {3, 1}, {3, 3}, {4, 2}, {4, 3}}

	grid := model.NewGrid(M, N)
	//grid.IntializeGrid(pts)
	grid.RandomInitialize()

	fmt.Println("intial Grid")
	grid.PrintGrid()
	for i := 0; i < ticks; i++ {
		nextGrid := service.UpdateMatrix(grid)
		grid.SetMatrix(nextGrid)
		fmt.Printf("After %d tick(s): \n", i+1)
		grid.PrintGrid()
	}

}
