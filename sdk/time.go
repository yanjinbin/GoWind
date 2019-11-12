package sdk

import (
	"time"
)

func Time() {

	format := "20060102150405"

	// t1 没有写 time.Now() 是为了避免秒以下单位的时间的影响
	// 除此之外和写 time.Now() 是一样的
	t1 := time.Date(2017, time.November, 30, 0, 0, 0, 0, time.Local)
	t2, _ := time.Parse(format, t1.Format(format))

	println("1-1 ", t1.Format(format) == t2.Format(format))
	println("1-2 ", t1 == t2)
	println("1-3 ", t1.Equal(t2))

	t2, _ = time.ParseInLocation(format, t1.Format(format), time.Local)
	println("2-1 ", t1.Format(format) == t2.Format(format))
	println("2-2 ", t1 == t2)
	println("2-3 ", t1.Equal(t2))

	t2 = t2.UTC()
	println("3-1 ", t1.Format(format) == t2.Format(format))
	println("3-2 ", t1 == t2)
	println("3-3 ", t1.Equal(t2))
}
