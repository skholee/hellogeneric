package main

type FilterElem interface {
	int | float64
}

// type parameter also can like: FilterByGeneric[V int | float64]
func FilterByGeneric[V FilterElem](elems []V, predicate func(V)bool) []V {
	var r []V
	for _, e := range elems {
		if predicate(e) {
			r = append(r, e)
		}
	}
	return r
}