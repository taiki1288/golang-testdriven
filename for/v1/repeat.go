package iteration

func Repeat(chatacter string) string {
	var repeated string 
	for i := 0; i < 5; i++ {
		repeated = repeated + chatacter
	}
	return repeated
}