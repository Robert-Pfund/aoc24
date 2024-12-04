package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/Robert-Pfund/aoc2024/day2/safety"
)

func main() {

	file, err := os.Open("./day2/input.txt")
	if err != nil {

		log.Fatalln("failed to open input file:", err)
	}
	defer file.Close()

	var lines [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		regex := regexp.MustCompile("[0-9]+")
		matches := (regex.FindAll([]byte(line), -1))

		var currentLine []int

		for idx, match := range matches {

			matchInt, err := strconv.Atoi(string(match))
			if err != nil {

				log.Fatalf("failed to convert string to int at index %d: %s", idx, err)
			}
			currentLine = append(currentLine, matchInt)
		}

		lines = append(lines, currentLine)
	}

	// Part 1 -
	safeReports := safety.GetSafeReports(lines)
	log.Println("number of safe reports:", safeReports)

}
