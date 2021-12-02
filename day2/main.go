package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/bAngerman/adventofcode/utils"
)

func calculateOffset(commands [][]string) (int, int, int) {
	aim := 0
	offsetX := 0
	offsetY := 0

	for _, command := range commands {
		dir := command[0]
		distance, _ := strconv.Atoi(command[1])

		if dir == "down" {
			// offsetY = offsetY + distance
			aim = aim + distance

		} else if dir == "up" {
			// offsetY = offsetY - distance
			aim = aim - distance

		} else if dir == "forward" {
			offsetX = offsetX + distance
			offsetY = offsetY + (aim * distance)
		}

		log.Println(dir, distance)
		log.Println("Current vals\t", offsetX, "\t", offsetY, "\t", aim)
		log.Println()
	}

	return offsetX, offsetY, aim
}

func main() {
	inputArr, _ := utils.ReadTextFile("day2/input")

	commands := make([][]string, 0)

	for _, str := range inputArr {
		split := strings.Fields(str)
		commands = append(commands, split)
	}

	offsetX, offsetY, _ := calculateOffset(commands)

	log.Println("Distance from start is:\t", "Horizontal", offsetX, "\t", "Vertical", offsetY)
	log.Println("Multiplied:", offsetX*offsetY)
}
