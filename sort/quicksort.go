package main

func quickSort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	v := s[0]
	var left, right []int
	for _, e := range s[1:] {
		if e <= v {
			left = append(left, e)
		} else {
			right = append(right, e)
		}
	}
	return append(append(quickSort(left), v), quickSort(right)...)
}

