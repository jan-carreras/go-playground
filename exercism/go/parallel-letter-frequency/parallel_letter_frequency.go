package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(list []string) FreqMap {
	rsp := FreqMap{}

	c := make(chan FreqMap, len(list))
	for _, s := range list {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}

	for range list {
		for k, v := range <-c {
			rsp[k] += v
		}
	}

	return rsp
}
