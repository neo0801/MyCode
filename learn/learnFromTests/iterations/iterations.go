package iterations

// LoopFor return target*n
func LoopFor(target string, times int) string {
	res := ""
	for i := 0; i < times; i++ {
		res += target
	}
	return res
}
