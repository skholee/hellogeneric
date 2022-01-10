package main

func FilterInts(elems []int, predicate func(int) bool) []int {
	var r []int
	for _, e := range elems {
		if predicate(e) {
			r = append(r, e)
		}
	}
	return r
}
