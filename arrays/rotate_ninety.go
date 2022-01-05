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
		for j := i; j < cubeSize-i-1; j++ {
			temp := cube[i][j]                 // 1
			cube[i][j] = cube[cubeSize-1-j][i] // 7 replaces 1
			printCube(cube)
			cube[cubeSize-1-j][i] = cube[cubeSize-1-i][cubeSize-1-j] // 9 replaces 7
			printCube(cube)
			cube[cubeSize-1-i][cubeSize-1-j] = cube[j][cubeSize-1-i] // 3 replaces 9
			printCube(cube)
			cube[j][cubeSize-1-i] = temp // 1 replaces 3
			printCube(cube)
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
