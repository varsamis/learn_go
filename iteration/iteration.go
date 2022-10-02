package iteration

func Repeat(character string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated

}

func Count(str string, substr string) int {
	var count int
	substrLength := len(substr)
	for i := 0; i <= len(str)-substrLength; i++ {
		if str[i:substrLength+i] == substr {
			count++
		}
	}
	return count
}
