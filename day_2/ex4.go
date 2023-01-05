package main

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

type calcFunc func(...float64) float64

func operation(op string) (calcFunc, bool) {
	switch op {
	case minimum:
		return minFunc, true
	case maximum:
		return maxFunc, true
	case average:
		return aveFunc, true
	}
	return nil, false
}

func minFunc(grades ...float64) float64 {
	min := grades[0]
	for _, v := range grades {
		if v < min {
			min = v
		}
	}
	return min
}

func maxFunc(grades ...float64) float64 {
	max := grades[0]
	for _, v := range grades {
		if v > max {
			max = v
		}
	}
	return max
}

func aveFunc(grades ...float64) float64 {
	var total float64
	for _, v := range grades {
		total += v
	}
	return total / float64(len(grades))
}
