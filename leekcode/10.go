package main

func isMatch(s string, p string) bool {
	if s == "" {
		return false
	}

	pIndex := 0
	lastIndex := -1
	for i := 0; i < len(s); i++ {

		switch string(p[pIndex]) {
		case `*`:
		case `.`:
		default:

		}
	}

}
