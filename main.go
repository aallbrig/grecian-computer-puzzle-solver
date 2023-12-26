package main

import "fmt"

// 5 layers, 4 rows, 12 elements in each row
// top layer to bottom layer, outer radial column to inner radial column
// -1 means no element
var grecianComputerPuzzle = [5][4][12]int{
	{
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{15, -1, 8, -1, 3, -1, 6, -1, 10, -1, 7, -1},
	},
	{
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{15, -1, -1, 14, -1, 9, -1, 12, -1, 4, -1, 7},
		{6, -1, 11, 11, 6, 11, -1, 6, 17, 7, 3, -1},
	},
	{
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{22, -1, 16, -1, 9, -1, 5, -1, 10, -1, 8, -1},
		{11, 26, 14, 1, 12, -1, 21, 8, 15, 4, 9, 18},
		{17, 4, 5, -1, 7, 8, 9, 13, 9, 7, 13, 21},
	},
	{
		{12, -1, 6, -1, 10, -1, 10, -1, 1, -1, 9, -1},
		{2, 13, 9, -1, 17, 19, 3, 12, 3, 26, 6, -1},
		{6, -1, 14, 12, 3, 8, 9, -1, 9, 20, 12, 3},
		{7, 14, 11, -1, 8, -1, 16, 2, 7, -1, 9, -1},
	},
	{
		{8, 3, 4, 12, 2, 5, 10, 7, 16, 8, 7, 8},
		{4, 4, 6, 6, 3, 3, 14, 14, 21, 21, 9, 9},
		{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		{11, 11, 14, 11, 14, 11, 14, 14, 11, 14, 11, 14},
	},
}

type GrecianComputer struct {
	Puzzle [5][4][12]int
}

func (gc *GrecianComputer) Solve() {
	rotationCounts := [5]int{}

	for a := 0; a < 12; a++ {
		for b := 0; b < 12; b++ {
			for c := 0; c < 12; c++ {
				for d := 0; d < 12; d++ {
					for e := 0; e < 12; e++ {
						if gc.IsSolved() {
							fmt.Println("Puzzle solved!")
							gc.PrintRotationCounts(rotationCounts)
							return
						}
						gc.RotateLayer(4)
						rotationCounts[4]++
					}
					gc.RotateLayer(3)
					rotationCounts[3]++
					rotationCounts[4] = 0
				}
				gc.RotateLayer(2)
				rotationCounts[2]++
				rotationCounts[3] = 0
			}
			gc.RotateLayer(1)
			rotationCounts[1]++
			rotationCounts[2] = 0
		}
		gc.RotateLayer(0)
		rotationCounts[0]++
		rotationCounts[1] = 0
	}

	fmt.Println("Puzzle not solved.")
	gc.PrintRotationCounts(rotationCounts)
}

// PrintRotationCounts prints the number of rotations for each layer.
func (gc *GrecianComputer) PrintRotationCounts(rotationCounts [5]int) {
	for i, count := range rotationCounts {
		fmt.Printf("Layer %d: %d rotations\n", i+1, count)
	}
}

func (gc *GrecianComputer) IsSolved() bool {
	for column := 0; column < 12; column++ {
		colSum := gc.ColumnSum(column)
		if colSum != 42 {
			return false
		}
	}
	return true
}
func (gc *GrecianComputer) RotateLayer(layerIndex int) {
	var layer = gc.Puzzle[layerIndex]
	var newLayer [4][12]int
	for row := 0; row < 4; row++ {
		for column := 0; column < 12; column++ {
			newLayer[row][(column+1)%12] = layer[row][column]
		}
	}
	gc.Puzzle[layerIndex] = newLayer
}
func (gc *GrecianComputer) ReadValue(row int, column int) int {
	for layer := 0; layer < len(gc.Puzzle); layer++ {
		if value := gc.Puzzle[layer][row][column]; value != -1 {
			return value
		}
	}
	return -1
}
func (gc *GrecianComputer) ColumnSum(column int) int {
	return gc.ReadValue(0, column) + gc.ReadValue(1, column) + gc.ReadValue(2, column) + gc.ReadValue(3, column)
}
func (gc *GrecianComputer) PrintAllLayers() {
	for _, layer := range gc.Puzzle {
		for _, row := range layer {
			for _, element := range row {
				if element == -1 {
					fmt.Printf("    ")
				} else {
					fmt.Printf("%3d ", element)
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
func (gc *GrecianComputer) Print() {
	for row := 0; row < 4; row++ {
		for column := 0; column < 12; column++ {
			if value := gc.ReadValue(row, column); value != -1 {
				fmt.Printf("%3d ", value)
			} else {
				fmt.Printf("    ")
			}
		}
		fmt.Println()
	}
	fmt.Println("================================================")

	for column := 0; column < 12; column++ {
		var columnSum = 0
		for row := 0; row < 4; row++ {
			if value := gc.ReadValue(row, column); value != -1 {
				columnSum += value
			}
		}
		fmt.Printf("%3d ", columnSum)
	}
	fmt.Println()
	fmt.Println()
}

func main() {
	var grecianComputer = GrecianComputer{
		Puzzle: grecianComputerPuzzle,
	}
	grecianComputer.Solve()
	grecianComputer.Print()
}
