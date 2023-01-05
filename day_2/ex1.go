package main

func calculateTaxes(salary float64) float64 {
	switch {
	case salary > 150000:
		{
			return salary * 0.27
		}
	case salary > 50000:
		{
			return salary * 0.17
		}
	default:
		{
			return 0
		}
	}
}
