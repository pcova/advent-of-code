package lens

func Hash(s string) int {
	var h int
	for _, c := range s {
		h = ((h + int(c)) * 17) % 256
	}
	return h
}
