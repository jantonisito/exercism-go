package letter

//import "sync"

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

// // ConcurrentFrequency counts the frequency of each rune in the given strings, by making use of concurrency.
// func ConcurrentFrequency(l []string) FreqMap {
// 	var wg sync.WaitGroup

// 	c := make(chan *FreqMap, len(l))
// 	fm := make([]*FreqMap, len(l))

// 	for i, s := range l {
// 		wg.Add(1)
// 		go func(s string) {
// 			defer wg.Done()
// 			x := Frequency(s)
// 			c <- &x
// 		}(s)
// 		// This causes the channel to block each iteration, which means that
// 		// the sending goroutines are blocked until a receiving goroutine reads
// 		// from the channel (see: ConcurrentFrequency2). Usage of WaitGroup in
// 		// this case is unnecessary since reading from channel will block anyway.
// 		fm[i] = <-c
// 	}
// 	wg.Wait()

// 	out := FreqMap{}
// 	for _, f := range fm {
// 		for k, v := range *f {
// 			out[k] += v
// 		}
// 	}
// 	return out
// }

// // ConcurrentFrequency counts the frequency of each rune in the given strings, by making use of concurrency.
// func ConcurrentFrequency(l []string) FreqMap {
// 	var wg sync.WaitGroup

// 	c := make(chan *FreqMap, len(l))

// 	for i := range l {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			x := Frequency(l[i])
// 			c <- &x
// 		}(i)
// 		// Note: we only write to channel - there is no blocking
// 	}
// 	// Blocking is done by WaitGroup - which is necessary here
// 	wg.Wait()

// 	out := make(FreqMap, 48)
// 	for range l {
// 		//we defer reading from channel until all work is done
// 		f := <-c
// 		for k, v := range *f {
// 			out[k] += v
// 		}
// 	}
// 	return out
// }
