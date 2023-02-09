package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	//m := FreqMap{}
	m := make(FreqMap, 48)
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings, by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {

	c := make(chan FreqMap, len(l))

	for i := range l {
		go func(i int) {
			c <- Frequency(l[i])
		}(i)
	}

	out := make(FreqMap, 48)
	for range l {
		//we defer reading from channel until all work is done
		f := <-c
		for k, v := range f {
			out[k] += v
		}
	}
	return out
}
