package main

import (
	"fmt"
	"math/rand"
	"time"
)

const rows int = 64
const cols int = 128

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UTC().UnixNano())

	// initialize world
	var world = make([][]rune, rows)
	for y := 0; y < rows; y++ {
		world[y] = make([]rune, cols)
		for x := 0; x < cols; x++ {
			world[y][x] = getRandomValue()
		}
	}

	// let the world life
	for {
		printWorld(world)
		world = evolveWorld(world)
		time.Sleep(time.Second / 2)
	}
}

func getRandomValue() rune {
    if rand.Intn(10) == 0 {
        return 'X'
    }

    return ' '
}

func printWorld(myworld [][]rune) {
	// clear screen
	fmt.Print("\033[H\033[2J")

	for y := 0; y < rows; y++ {
		fmt.Println(string(myworld[y]))
	}
}

func evolveWorld(myworld [][]rune) [][]rune {
	var newworld = make([][]rune, rows)
	for y := 0; y < rows; y++ {
		newworld[y] = make([]rune, cols)
		for x := 0; x < cols; x++ {
			var neighborCount = getNumOfNeighbors(myworld, x, y)
			var current = myworld[y][x]
			var newVal = current

			if current == ' ' && neighborCount == 3 {
				newVal = 'X'
			} else if current == 'X' && (neighborCount < 2 || neighborCount > 3) {
				newVal = ' '
			}

			newworld[y][x] = newVal
		}
	}

	return newworld
}

func getNumOfNeighbors(myworld [][]rune, x int, y int) int {
	var num int = 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			var currentX = x + i
			var currentY = y + j

			if currentX >= 0 && currentX < cols && currentY >= 0 && currentY < rows && (i != 0 || j != 0) && myworld[currentY][currentX] == 'X' {
				num++
			}
		}
	}

	return num
}
