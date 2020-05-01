package leetcode


// LC 57
func insert(intervals [][]int, newInterval []int) [][]int {
	s := newInterval[0]
	e := newInterval[1]
	l := make([][]int, 0)
	r := make([][]int, 0)
	for _, v := range intervals {
		if v[1] < s {
			l = append(l, v)
		} else if v[0] > e {
			r = append(r, v)
		} else {
			s = min(s, v[0])
			e = max(e, v[1])
		}
	}
	ans := make([][]int, 0)
	ans = append(ans, l...)
	ans = append(ans, []int{s, e})
	ans = append(ans, r...)
	return ans

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
