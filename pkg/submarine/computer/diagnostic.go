package computer

import (
	"fmt"
)

func PowerConsumption(reports []int, bitRate int) int {
	reportLength := len(reports)
	analysis := reportAnalysis(reports, bitRate)
	fmt.Printf("analysis: %v\n", analysis)
	gamma := powerGammaRate(analysis, reportLength)
	fmt.Printf("gamma: %b\n", gamma)
	epsilon := powerEpsilonRate(analysis, reportLength)
	fmt.Printf("epsilon: %b\n", epsilon)
	return gamma * epsilon
}

func LifeSupportRating(reports []int, bitRate int) int {
	oxygen := oxygenRate(reports, bitRate)
	fmt.Printf("oxygen: %b\n", oxygen)
	co2 := co2ScrubbingRate(reports, bitRate)
	fmt.Printf("co2: %b\n", oxygen)
	return oxygen * co2
}

func oxygenRate(reports []int, bitRate int) int {
	position := bitRate - 1
	commonBit := 0
	onesCount := 0

	for len(reports) > 1 {
		fmt.Printf("---------------------------\n")
		onesCount = numberOfOnes(reports, position)
		fmt.Printf("oxygen position: %v\n", position)
		fmt.Printf("oxygen ones Count: %v\n", onesCount)
		commonBit = mostCommonNumber(onesCount, len(reports), true)
		fmt.Printf("oxygen common bit: %v\n", commonBit)
		reports = similarReports(reports, commonBit, position)
		fmt.Printf("reports %b, %v\n", commonBit, position)
		fmt.Printf("reports %b\n", reports)
		position--
	}

	return reports[0]
}

func co2ScrubbingRate(reports []int, bitRate int) int {
	position := bitRate - 1
	commonBit := 0
	onesCount := 0

	for len(reports) > 1 {
		fmt.Printf("---------------------------\n")
		onesCount = numberOfOnes(reports, position)
		fmt.Printf("co2 position: %v\n", position)
		fmt.Printf("co2 ones Count: %v\n", onesCount)
		commonBit = leastCommonNumber(onesCount, len(reports), false)
		fmt.Printf("co2 common bit: %v\n", commonBit)
		reports = similarReports(reports, commonBit, position)
		fmt.Printf("co2 %b, %v\n", commonBit, position)
		fmt.Printf("co2 %b\n", reports)
		position--
	}

	return reports[0]
}

func similarReports(reports []int, number int, position int) []int {
	var result []int
	for _, report := range reports {
		if GetBit(report, position) == number {
			result = append(result, report)
		}
	}
	return result
}

func reportAnalysis(reports []int, bitRate int) []int {
	analysis := make([]int, bitRate)
	for i, _ := range analysis {
		analysis[i] = numberOfOnes(reports, i)
	}
	return analysis
}

func numberOfOnes(reports []int, position int) int {
	count := 0
	for _, report := range reports {
		if HasOne(report, position) {
			count++
		}
	}
	return count
}

func mostCommonNumber(onesCount int, total int, preferOnes bool) int {
	if onesCount == total-onesCount {
		if preferOnes {
			return 1
		} else {
			return 0
		}
	} else {
		if onesCount > total-onesCount {
			return 1
		} else {
			return 0
		}
	}
}

func leastCommonNumber(onesCount int, total int, preferOnes bool) int {
	if onesCount == total-onesCount {
		if preferOnes {
			return 1
		} else {
			return 0
		}
	} else {
		if onesCount > total-onesCount {
			return 0
		} else {
			return 1
		}
	}
}

func powerGammaRate(analysis []int, total int) int {
	rate := 0
	for i, onesCount := range analysis {
		rate = UpdateBit(rate, i, mostCommonNumber(onesCount, total, false))
	}
	return rate
}

func powerEpsilonRate(analysis []int, total int) int {
	rate := 0
	for i, onesCount := range analysis {
		rate = UpdateBit(rate, i, leastCommonNumber(onesCount, total, true))
	}
	return rate
}
