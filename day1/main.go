package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/Robert-Pfund/aoc2024/day1/distance"
	"github.com/Robert-Pfund/aoc2024/day1/similarity"
)

func main() {

	file, err := os.Open("./day1/input.txt")
	if err != nil {

		log.Fatalln("failed to open input file:", err)
	}
	defer file.Close()

	var left []int
	var right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		regex := regexp.MustCompile("[0-9]+")
		numbers := (regex.FindAll([]byte(line), -1))

		x, err := strconv.Atoi(string(numbers[0]))
		if err != nil {

			log.Fatalln("failed to convert number 1 from string to int:", err)
		}
		left = append(left, x)
		y, err := strconv.Atoi(string(numbers[1]))
		if err != nil {

			log.Fatalln("failed to convert number 2 from string to int:", err)
		}
		right = append(right, y)
	}

	// Part 1 -
	distance := distance.GetDistance(left, right)
	log.Println("distance:", distance)

	// Part 2 -
	similarity := similarity.GetSimilarity(left, right)
	log.Println("similarity:", similarity)
}
