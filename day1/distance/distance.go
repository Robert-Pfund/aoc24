package distance

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func GetDistance() {

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

	slices.Sort(left)
	slices.Sort(right)

	if len(left) != len(right) {

		log.Fatalln("got lists with unequal lengths")
	}

	distanceSum := 0
	for idx, x := range left {

		y := right[idx]
		difference := x - y
		if difference < 0 {
			difference = difference * -1
		}

		distanceSum += difference
		log.Println(distanceSum)
	}
}
