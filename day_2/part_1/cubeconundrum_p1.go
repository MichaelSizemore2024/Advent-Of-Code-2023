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

	var total int
	var valid bool = true
	// Loops through lines in file
	for fileScanner.Scan() {
		// Gets current line
		line := fileScanner.Text()

		// Spilts game # and gameValues
		gameValues := strings.Split(line, ":")

		//Gets set number
		gameNumber := strings.Split(gameValues[0], " ")

		//Spilts up by games seperated by ';'
		games := strings.Split(gameValues[1], ";")

		// Sets a game as valid by default
		valid = true

		// Loops through each game in the current set
		for i := range games {
			// Gets the values for each game
			gameValues = strings.Split(games[i], ",")
			// Loops through each and checks if valid
			for j := range gameValues {
				colorValue := strings.Split(gameValues[j], " ")
				blocks, err := strconv.Atoi(colorValue[1])
				if err != nil {
					fmt.Println("Error converting string to integer:", err)
					return
				}
				switch colorValue[2] {
				case "red":
					if blocks > 12 {
						valid = false
						break
					}
				case "green":
					if blocks > 13 {
						valid = false
						break
					}
				case "blue":
					if blocks > 14 {
						valid = false
						break
					}
				}
			}
		}
		// Checks if current set passed all the tests
		if valid {
			// Convert string to integer
			intNumber, err := strconv.Atoi(gameNumber[1])
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return
			}
			total += intNumber
		}
	}

	// Closes File
	readFile.Close()

	// Prints total
	fmt.Println("The sum of valid game ids is:", total)
}
