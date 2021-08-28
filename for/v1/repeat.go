package iteration

const repeatConst = 5

func Repeat(character string) string {
	var repeated string 
	for i := 0; i < repeatConst; i++ {
		repeated += character
	}
	return repeated
}