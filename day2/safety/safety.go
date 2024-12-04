package safety

import "math"

func GetSafeReports(list [][]int) int {

	safeReports := 0
	var safe bool

	for _, reports := range list {

		safe = true
		trend := reports[1] > reports[0]

		for idx := 1; idx < len(reports); idx++ {

			if idx == 0 {
				break
			}

			if reports[idx] == reports[idx-1] {
				safe = false
				break
			}

			if (reports[idx] > reports[idx-1]) != trend {
				safe = false
				break
			}

			if math.Abs(float64(reports[idx])-float64(reports[idx-1])) > 3 {
				safe = false
				break
			}
		}

		if safe {
			safeReports++
		}
	}

	return safeReports
}
