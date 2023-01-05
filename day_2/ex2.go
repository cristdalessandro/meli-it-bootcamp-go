package main

func gradesAverage(grades ...float64) float64 {
	var total float64
	for _, v := range grades {
		if v < 0 {
			panic("No puede haber notas negativas")
		}
		total += v
	}
	return total / float64(len(grades))
}
