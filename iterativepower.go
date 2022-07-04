package piscine

func IterativePower(nb int, power int) int {
	a := nb
	if nb == 0 || power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else if power == 1 {
		return nb
	} else {
		for i := 1; i < power; i++ {
			nb = nb * a
		}
		return nb
	}
}
