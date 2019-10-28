func hIndex(citations []int) int {
    l := len(citations)
    sort.Ints(citations)
	for i := l - 1; i >= 0; i-- {
		if citations[l-1-i] >= i+1 {
			return i + 1
		}
	}
	return 0
}
