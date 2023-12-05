package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file for reading
	readFile, err := os.Open("games.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
	}

	// Create a scanner to read file
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// Initializes needed varibles
	var total, red, blue, green int

	// Loops through lines in file
	for fileScanner.Scan() {
		// Gets current line
		line := fileScanner.Text()

		// Spilts game # and gameValues
		gameValues := strings.Split(line, ":")

		//Spilts up by games seperated by ';'
		games := strings.Split(gameValues[1], ";")

		// Resets values
		red = 0
		blue = 0
		green = 0

		for i := range games {
			// Gets the values for each game
			gameValues = strings.Split(games[i], ",")

			// Loops through each
			for j := range gameValues {
				colorValue := strings.Split(gameValues[j], " ")
				totalBlocks, err := strconv.Atoi(colorValue[1])
				if err != nil {
					fmt.Println("Error converting string to integer:", err)
					return
				}
				switch colorValue[2] {
				case "red":
					if totalBlocks > red {
						red = totalBlocks
					}
				case "green":
					if totalBlocks > green {
						green = totalBlocks
					}
				case "blue":
					if totalBlocks > blue {
						blue = totalBlocks
					}
				}
			}
		}
		total += red * blue * green
	}

	// Closes File
	readFile.Close()

	// Prints total
	fmt.Println("The sumof the power of these sets is:", total)
}
