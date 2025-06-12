package reverse

// Reverse reverses string param
func Reverse(input string) string {
	return reverse1(input)
}

func reverse1(input string) string {

	runes := []rune(input)
	length := len(runes)

	//Benchmark:
	//BenchmarkReverse-8
	//4447509               274.3 ns/op             0 B/op          0 allocs/op
	// Note: i < length/2 - due to 0 indexing
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)
}

func reverse2(input string) string {

	if len(input) == 0 {
		return input
	}

	runes := []rune(input)
	length := len(runes)
	reverse := make([]rune, length)

	//Benchmark:
	//BenchmarkReverse-8
	//1895503               534.9 ns/op           152 B/op          5 allocs/op
	for i, ch := range runes {
		reverse[length-1-i] = ch
	}
	return string(reverse)
}
