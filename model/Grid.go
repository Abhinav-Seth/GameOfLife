package model

import (
	"fmt"
	"math/rand"
)

type Grid struct {
	m, n   int
	matrix *[][]uint8
}

// add nil check for getter
func (g *Grid) GetM() int {
	return g.m
}

func (g *Grid) SetM(m int) {
	g.m = m
}

func (g *Grid) GetN() int {
	return g.n
}

func (g *Grid) SetN(n int) {
	g.n = n
}

func (g *Grid) GetMatrix() *[][]uint8 {
	return g.matrix
}

func (g *Grid) SetMatrix(matrix *[][]uint8) {
	g.matrix = matrix
}

// what will happen if matrix is not a pointer
func NewGrid(m int, n int) *Grid {
	grid := &Grid{m: m, n: n}

	grid.matrix = makeArray(m, n)
	return grid
}

func makeArray(m int, n int) *[][]uint8 {
	arr := make([][]uint8, m)
	for i := range arr {
		arr[i] = make([]uint8, n)
	}
	return &arr
}

// better way to intialize
func (g *Grid) RandomInitialize() {
	m := g.m
	n := g.n
	arr := *g.matrix
	for k := 0; k < rand.Intn(m*n); k++ {
		i := rand.Intn(m)
		j := rand.Intn(n)
		arr[i][j] = 1
	}
}

func (g *Grid) IntializeGrid(positions []Position) {
	arr := *g.matrix
	for _, pt := range positions {
		i := pt.GetX()
		j := pt.GetY()
		arr[i][j] = 1
	}

}

func (g *Grid) PrintGrid() {
	mat := *(g.matrix)
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 0 {
				fmt.Printf(" . ")
			} else {
				fmt.Printf(" | ")
			}
		}
		fmt.Println()
	}
}
