package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Converts input into single digit, or returns what was given
func wordToNumber(word string) string {
	numberMap := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":  "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	if num, found := numberMap[word]; found {
		return num
	}

	return word
}

func main() {
	// Open the file for reading
	readFile, err := os.Open("calibration.txt")
    if err != nil {
        fmt.Println("Error opening the file:", err)
    }

    // Create a scanner to read file
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
  
    // Sets up a Regex to check if digit is in provided string
	re := regexp.MustCompile("[0-9]|one|two|three|four|five|six|seven|eight|nine")
	
	// Sets up ongoing total
	var total int

	// Loops through lines in file
    for fileScanner.Scan() {
		line := fileScanner.Text()

		// Loops through line incrementally till it finds a match
		var first string = ""
		for _, char := range line {
			first += string(char)
			var arr = re.FindAllString(first, -1)
			// Saves the first value when a match is found
			if len(arr) != 0 {
				first = wordToNumber(arr[0])
				break
			} 
		}

		// Loops through line backwards incrementally till it finds a match
		var second string = ""
		for i := len(line) - 1; i >= 0; i-- {
			second = string(line[i]) + second
			var arr = re.FindAllString(second, -1)
			// Adds the new value onto the previous one
			if len(arr) != 0 {
				first += wordToNumber(arr[0])
				break
			}  
		}

		// Casts resulting value to an int
		num, err :=strconv.Atoi(first)
		if err != nil {
			fmt.Println("Error casting string to an int:", err)
		}

		// Adds to ongoing total
		total += num
    }

	// Closes File
    readFile.Close()

	// Prints total
	fmt.Println("The sum of calibration values is:", total)
}