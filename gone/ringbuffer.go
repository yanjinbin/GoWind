package gone

// 初版  问题出在 full 和 empty怎么判定呢 http://bit.ly/2KiKjlr
var Read int
var Write int
var array [100]int

func mask(idx int) int {
	return idx & (cap(array))
}

func incr(idx int) int {
	return mask(idx + 1)
}

func push(val int) {
	array[mask(Write)] = val
	Write++
}

func shift() int {
	ret := array[Read]
	Read++
	return ret
}

func empty() bool {
	return Read == Write
}

func full() bool {
	return size() == cap(array)
}
func size() int {
	return Write - Read
}

// 初版  人为的把 read write 取模更新了  失败之笔啊  read write本来就是单调递增的

func push_02(val int) {
	array[Write] = val
	Write = incr(Write)
}

func shift_02() int {
	ret := array[Read]
	Read = incr(Read)
	return ret
}

func empty_02() bool {
	return Read == Write
}

func size_02() int {
	return mask(Write - Read)
}

// 改进  去掉read 换 length
var length int

func push_1(val int) {
	array[mask(Read+length)] = val
	length++
}

func shift_1() int {
	length--
	ret := array[mask(Read)]
	incr(Read)
	return ret
}

func Empty_1() bool {
	return length == 0
}

func size_1() int {
	return length
}

// 最终版
