package main

func FilterFloat64s(elems []float64, predicate func(float64) bool) []float64 {
	var r []float64
	for _, e := range elems {
		if predicate(e) {
			r = append(r, e)
		}
	}
	return r
}
