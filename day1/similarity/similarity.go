package similarity

func GetSimilarity(left, right []int) int {

	similarity := 0

	for _, leftValue := range left {

		count := 0
		for _, rightVale := range right {

			if leftValue == rightVale {

				count++
			}
		}

		similarity += leftValue * count
	}

	return similarity
}
