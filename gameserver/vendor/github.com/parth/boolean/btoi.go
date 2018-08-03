package boolean

func BtoI(b bool) int {
	if b {
		return 1
	}

	return 0
}

func ItoB(i int) bool {
	return i != 0
}
