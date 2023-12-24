package iteration

func Repeat(char string, repeat int) (RepeatedChar string) {
	for i := 0; i < repeat; i++ {
		RepeatedChar += char
	}
	return
}
