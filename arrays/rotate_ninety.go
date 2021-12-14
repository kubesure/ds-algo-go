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

	for i := 0; i < cubeSize/2; i++ {
		for j := i; j < cubeSize-1; j++ {
			temp := cube[i][j]
			cube[i][j] = cube[cubeSize-1-j][i] // 7 > 1 // 4 > 2
			//printCube(cube)
			cube[cubeSize-1-j][i] = cube[cubeSize-1][cubeSize-1-j] // 9 > 7 // 8 > 4
			//printCube(cube)
			cube[cubeSize-1][cubeSize-1-j] = cube[j-i][cubeSize-1] // 3 > 9 // 6 > 8
			//printCube(cube)
			cube[j-i][cubeSize-1] = temp // 1 > 3 // 8 > 4
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
		fmt.Printf("\n \n")
	}
	fmt.Printf("\n")
}
