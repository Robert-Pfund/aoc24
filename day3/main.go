package main

import (
	"bufio"
	"log"
	"os"
	"regexp"

	"github.com/Robert-Pfund/aoc2024/day3/multiplication"
)

func main() {

	file, err := os.Open("./day3/input.txt")
	if err != nil {

		log.Fatalln("failed to open input file:", err)
	}
	defer file.Close()

	var allMultiplications []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		regex := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
		multiplications := (regex.FindAll([]byte(line), -1))

		for _, multiplication := range multiplications {

			allMultiplications = append(allMultiplications, string(multiplication))
		}
	}

	// Part 1 -
	sum := multiplication.GetMultiplication(allMultiplications)
	log.Println("sum of all multiplications:", sum)
}
