package scrape

// ByLength is a string interface implementaion to be used for sorting
type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	if len(s[i]) == len(s[j]) {
		return s[i] > s[j]
	}
	return len(s[i]) > len(s[j])
}
