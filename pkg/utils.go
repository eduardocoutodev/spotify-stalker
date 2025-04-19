package pkg

func FindIndex[S ~[]E, E comparable](s S, filter func(E) bool) int {
	for i := range s {
		if filter(s[i]) {
			return i
		}
	}
	return -1
}
