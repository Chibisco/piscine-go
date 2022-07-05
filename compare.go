package piscine

func Compare(a, b string) int {
	sentence1 := []rune(a)
	sentence2 := []rune(b)

	for i := 0; i < len(a)+len(b); i++ {
		if len(a) == len(b) && sentence1[i] == sentence2[i] {
			return 0
		} else if len(a) != len(b) && sentence1[0] != sentence2[0] {
			return -1
		} else if len(a) != len(b) && sentence1[len(a)-1] != sentence2[len(b)-1] {
			return 1
		}
	}
	return 0
}
