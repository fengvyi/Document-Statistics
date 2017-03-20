package lib
// Histogram ...
type Histogram map[string]int
type Character map[byte]int


// RemoveMax element with biggest value from histogram.
func (h Histogram) RemoveMax() (string,int) {
	max,count := h.maxByValue()
	delete(h, max)
	return max,count
}


func (h Histogram) maxByValue() (string,int) {
	var (
		maxKey string
		maxVal int
	)

	for key, val := range h {
		if val > maxVal {
			maxKey = key
			maxVal = val
		}
	}
	return maxKey, maxVal
}

