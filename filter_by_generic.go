package main

func FilterByGeneric[V int | float64](elems []V, predicate func(V)bool) []V {
	var r []V
	for _, e := range elems {
		if predicate(e) {
			r = append(r, e)
		}
	}
	return r
}