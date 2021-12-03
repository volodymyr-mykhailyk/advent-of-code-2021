package computer

func PowerConsumption(reports []string) int {
	reportLength := len(reports)
	analysis := reportAnalysis(reports)
	gamma := gammaRate(analysis, reportLength)
	epsilon := epsilonRate(analysis, reportLength)
	return gamma * epsilon
}

func reportAnalysis(reports []string) []int {
	analysis := make([]int, len(reports[0]))
	for _, report := range reports {
		for i, _ := range analysis {
			if report[i] == '1' {
				analysis[i]++
			}
		}
	}
	return analysis
}

func gammaRate(analysis []int, total int) int {
	length := len(analysis)
	rate := 0
	for i, onesCount := range analysis {
		if onesCount > total/2 {
			rate = SetBit(rate, length - i -1)
		}
	}
	return rate
}

func epsilonRate(analysis []int, total int) int {
	length := len(analysis)
	rate := 0
	for i, onesCount := range analysis {
		if onesCount < total/2 {
			rate = SetBit(rate, length - i - 1)
		}
	}
	return rate
}
