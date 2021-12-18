package main

import (
	"errors"
	"log"
	"strings"
	"unicode"

	"github.com/bAngerman/adventofcode/utils"
	"github.com/kr/pretty"
)

type room struct {
	name      string
	isBigRoom bool

	// Pointer for next room idx(s) and previous room idx
	prevRoom  int
	nextRooms []int
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

func pathfind(roomCopy []room, curRoomIdx, pathCount int) int {

	// curRoom := roomCopy[curRoomIdx]
	// prevRoom := roomCopy[curRoom.prevRoom]

	// pretty.Println("Currently in", curRoom.name)

	// // Room has been visited already, and it's a small room (only 1 visit allowed)
	// // if curRoom.visited && !curRoom.isBigRoom {
	// // 	pretty.Println("We have already visited ", curRoom.name, "so this path is invalid.")

	// // 	return 0
	// // }

	// if len(curRoom.nextRooms) != 0 {
	// 	for _, nextRoomIdx := range curRoom.nextRooms {
	// 		// nextRoom := roomCopy[nextRoomIdx]

	// 		// return pathfind(roomCopy, nextRoomIdx, pathCount)
	// 		// if len(nextRoom.nextRooms) != 0 {
	// 		// } else {

	// 		// 	pretty.Println("There are no rooms after the next room", nextRoom.name)
	// 		// }
	// 	}
	// } else {
	// 	pretty.Println("There are no next rooms connected to", curRoom.name)
	// 	// Made it to the end, incrementing path count.
	// 	if curRoom.name == "end" {
	// 		pretty.Println("made it to end")
	// 		return pathCount + 1
	// 	}

	// 	// if !prevRoom.isBigRoom {
	// 	// 	pretty.Println("The previous room is")
	// 	// 	return 0
	// 	// }
	// }

	return pathCount
}

func part1(rooms []room) {
	var distinctPathCount int

	pretty.Println("Rooms", rooms)

	startRoomIdx, err := findRoomByName(rooms, "start")
	if err != nil {
		log.Panic(err)
	}

	startRoom := rooms[startRoomIdx]
	for _, roomIdx := range startRoom.nextRooms {
		roomCopy := rooms
		distinctPathCount = distinctPathCount + pathfind(roomCopy, roomIdx, 0)
	}

	// pretty.Println(rooms)
	pretty.Println("Count of distinct paths", distinctPathCount)
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

				idx := 0
				found := false

				for nextRoomIdx, nextRoom := range rooms {
					if nextRoom.name == parts[1] {
						idx = nextRoomIdx
						found = true
					}
				}

				if !found {
					panic("Something went wrong when finding the next room index Sadgeg")
				}

				rooms[roomIdx].nextRooms = append(rooms[roomIdx].nextRooms, idx)
			}
		}
	}

	part1(rooms)
	part2(rooms)
}
