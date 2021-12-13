package arrays

import "fmt"

func rotate90(cube [][]int) ([][]int, error) {
	rows := len(cube)
	columns := len(cube[0])
	cubeSize := len(cube)

	if rows == 0 || rows != columns {
		return nil, fmt.Errorf("not a cube")
	}

	printCube(cube)

	for layer := 0; layer < cubeSize/2; layer++ {
		for j := layer; j < cubeSize-1-layer; j++ {
			temp := cube[layer][j]
			cube[layer][j] = cube[cubeSize-1-j][layer] // 7 > 1 // 4 > 2
			//printCube(cube)
			cube[cubeSize-1-j][layer] = cube[cubeSize-1][cubeSize-1-j] // 9 > 7 // 8 > 4
			//printCube(cube)
			cube[cubeSize-1][cubeSize-1-j] = cube[layer+j][cubeSize-1] // 3 > 9 // 6 > 8
			//printCube(cube)
			cube[layer+j][cubeSize-1] = temp // 1 > 3 // 8 > 4
			//printCube(cube)
		}
	}

	return cube, nil
}

func printCube(cube [][]int) {
	for _, row := range cube {
		for _, column := range row {
			fmt.Printf("%v ", column)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
