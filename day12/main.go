package main

import (
	"errors"
	"strings"
	"unicode"

	"github.com/bAngerman/adventofcode/utils"
	"github.com/kr/pretty"
)

type room struct {
	name      string
	isBigRoom bool

	// Pointer for next room idx(s) and previous room idx
	nextRooms []string
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func findRoomByName(rooms []room, name string) (int, error) {
	for idx, room := range rooms {
		if room.name == name {
			return idx, nil
		}
	}

	return 0, errors.New("Unable to find room with name " + name)
}

func part1(rooms []room) {
	var pathCount int
	pretty.Println(rooms)

	/**
	 * Starting at the "start" room, start building arrays of
	 * room names (paths) by going through each next connected room.
	 * Need to do some recursion to continue finding next rooms.
	 */

	start, _ := findRoomByName(rooms, "start")

	for _, nextRoom := range rooms[start].nextRooms {
	}

}

func part2(rooms []room) {

}

func main() {
	inputArr, _ := utils.ReadTextFile("day12/input2")

	var rooms []room

	// Create our rooms slice with just the names for now.
	for _, line := range inputArr {
		foundThis := false
		foundNext := false
		parts := strings.Split(line, "-")

		thisRoom := room{name: parts[0], isBigRoom: IsUpper(parts[0])}
		nextRoom := room{name: parts[1], isBigRoom: IsUpper(parts[1])}

		for _, room := range rooms {
			if room.name == thisRoom.name {
				foundThis = true
			}

			if room.name == nextRoom.name {
				foundNext = true
			}
		}

		if !foundThis {
			rooms = append(rooms, thisRoom)
		}

		if !foundNext {
			rooms = append(rooms, nextRoom)
		}
	}

	// Now build the connections to next rooms.
	for _, line := range inputArr {
		parts := strings.Split(line, "-")

		for roomIdx, room := range rooms {
			if room.name == parts[0] {

				name := ""

				for _, nextRoom := range rooms {
					if nextRoom.name == parts[1] {
						name = nextRoom.name
					}
				}

				rooms[roomIdx].nextRooms = append(rooms[roomIdx].nextRooms, name)
			}
		}
	}

	part1(rooms)
	part2(rooms)
}
