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
	var world [rows][cols]int

	rand.Seed(time.Now().UTC().UnixNano())

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if rand.Intn(10) == 0 {
				world[y][x] = 1
			} else {
				world[y][x] = 0
			}
		}
	}

	for {
		printWorld(world)
		world = evolveWorld(world)
		time.Sleep(time.Second / 2)
	}
}

func printWorld(myworld [rows][cols]int) {
	clearScreen()

	var output string
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if myworld[y][x] == 0 {
				output = output + " "
			} else {
				output = output + "X" 
			}
		}
		fmt.Println(output)
		output = ""
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func evolveWorld(myworld [rows][cols]int) [rows][cols]int {
	var newworld [rows][cols]int

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			var neighbors = checkNeighbors(myworld, x, y)
			var current = myworld[y][x]
			var newval = current

			if current == 0 && neighbors == 3 {
				newval = 1
			} else if current == 1 {
				if neighbors < 2 || neighbors > 3 {
					newval = 0
				}
			}

			newworld[y][x] = newval
		}
	}
	return newworld
}

func checkNeighbors(myworld [rows][cols]int, x int, y int) int {
	var num int = 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			var aktx = x + i
			var akty = y + j

			if aktx >= 0 && aktx < cols && akty >= 0 && akty < rows {
				if i != 0 || j != 0 {
					if myworld[akty][aktx] == 1 {
						num++
					}
				}
			}
		}
	}

	return num
}
