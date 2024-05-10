package iteration

func Repeat(char string, amnt int) string {
	var repeated string
	for i := 0; i < amnt; i++ {
		repeated += char
	}
	return repeated
}
