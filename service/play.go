package service

import (
	"GameOfLife/model"
	"GameOfLife/util"
)

func UpdateMatrix(grid *model.Grid) *[][]uint8 {

	var finalMatrix [][]uint8
	intialMatrix := *grid.GetMatrix()
	for i := range intialMatrix {
		var subMatrix []uint8
		for j := range intialMatrix[i] {
			isLive := intialMatrix[i][j] == 1
			liveN := countLiveNeighbors(intialMatrix, i, j)
			if (isLive && (liveN == 2 || liveN == 3)) || (!isLive && liveN == 3) {
				subMatrix = append(subMatrix, 1)
			} else {
				subMatrix = append(subMatrix, 0)
			}
		}
		finalMatrix = append(finalMatrix, subMatrix)
	}
	return &finalMatrix
}

func countLiveNeighbors(initialGrid [][]uint8, r int, c int) int {
	countLiveN := 0
	M := len(initialGrid)
	if M > 0 {
		N := len(initialGrid[0])
		for i := util.Max(r-1, 0); i <= util.Min(r+1, M-1); i++ {
			for j := util.Max(c-1, 0); j <= util.Min(c+1, N-1); j++ {
				if initialGrid[i][j] == 1 {
					countLiveN++
				}
			}
		}
		if initialGrid[r][c] == 1 {
			countLiveN--
		}
	}
	return countLiveN
}
