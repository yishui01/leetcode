func superPow(a int, b []int) int {
    const N = 1337
	remains := []int{}
	remain := a % N
	if remain == 0 {
		return 0
	}
	t := remain
	remains = append(remains, t)
	for {
		t = (t * remain) % N
		if t == remains[0] {
			break
		}
		remains = append(remains, t)
	}
	k := 0
	for i, v := range b {
		if i == len(b)-1 {
			k = (k*10 + v - 1) % len(remains)
		} else {
			k = (k*10 + v) % len(remains)
		}
	}
	return remains[k]
}
