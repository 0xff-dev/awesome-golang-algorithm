package Solution

func Solution(n int) bool {
	aliceWin := false
	removes := 10
	for n >= removes {
		n -= removes
		removes--
		aliceWin = !aliceWin
	}
	return aliceWin
}
