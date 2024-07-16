package math

func AddFloats(xc []float64) (sum float64) {
	for _, v := range xc {
		sum += v
	}
	return sum / float64(len(xc))
}
