package TFilter

func maxChannels(x int) int {
	for i := 10; i > 0; i-- {
		if x%i == 0 {
			return i
		}
	}
	return 1
}
