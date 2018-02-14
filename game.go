package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const rows int = 64
const cols int = 128

func main() {
	var world = make([][]rune, rows)

	rand.Seed(time.Now().UTC().UnixNano())

	for y := 0; y < rows; y++ {
		world[y] = make([]rune, cols)
		for x := 0; x < cols; x++ {
			if rand.Intn(10) == 0 {
				world[y][x] = 'X'
			} else {
				world[y][x] = ' '
			}
		}
	}

	for {
		printWorld(world)
		world = evolveWorld(world)
		time.Sleep(time.Second / 2)
	}
}

func printWorld(myworld [][]rune) {
	clearScreen()

	for y := 0; y < rows; y++ {
		fmt.Println(string(myworld[y]))
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func evolveWorld(myworld [][]rune) [][]rune {
	var newworld = make([][]rune, rows)

	for y := 0; y < rows; y++ {
		newworld[y] = make([]rune, cols)
		for x := 0; x < cols; x++ {
			var neighbors = checkNeighbors(myworld, x, y)
			var current = myworld[y][x]
			var newval = current

			if current == ' ' && neighbors == 3 {
				newval = 'X'
			} else if current == 'X' {
				if neighbors < 2 || neighbors > 3 {
					newval = ' '
				}
			}

			newworld[y][x] = newval
		}
	}

	return newworld
}

func checkNeighbors(myworld [][]rune, x int, y int) int {
	var num int = 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			var aktx = x + i
			var akty = y + j

			if aktx >= 0 && aktx < cols && akty >= 0 && akty < rows && (i != 0 || j != 0) && myworld[akty][aktx] == 'X' {
				num++
			}
		}
	}

	return num
}
