package multiplication

import (
	"log"
	"regexp"
	"strconv"
)

func GetMultiplication(statements []string) int {

	sum := 0

	for _, statement := range statements {

		regex := regexp.MustCompile(`[0-9]+`)
		factors := (regex.FindAll([]byte(statement), -1))

		factor1, err := strconv.Atoi(string(factors[0]))
		if err != nil {

			log.Fatalln("failed to convert to int (factor1):", err)
		}
		factor2, err := strconv.Atoi(string(factors[1]))
		if err != nil {

			log.Fatalln("failed to convert to int (factor2):", err)
		}

		product := factor1 * factor2
		sum += product
	}

	return sum
}
