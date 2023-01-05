package main

func salaryCalculation(minutes int, cat rune) float64 {
	hrs := (float64(minutes) / 60)
	switch {
	case cat == 'A':
		{
			return float64(1000) * hrs
		}
	case cat == 'B':
		{
			base := float64(1500) * hrs
			return base * 1.2
		}
	case cat == 'C':
		{
			base := float64(3000) * hrs
			return base * 1.5
		}
	}
	panic("categoria invalida")
}
