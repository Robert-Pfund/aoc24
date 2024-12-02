package distance

import (
	"log"
	"slices"
)

func GetDistance(left, right []int) int {

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

	return distanceSum
}
