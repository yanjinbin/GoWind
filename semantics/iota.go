package semantics

import "fmt"

// https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3

type Weekday int

// How to use iota?
// Iota expression is repeated by the other constants until another assignment or type declaration shows up.
const (
	Monday = iota // 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
	_         // 8
	NotDefine // 9
)

/*const (
	One   = 1
	Two   = 2
	// Three = 2 + 1 => 3
	// iota in the middle
	Three = iota + 1
	// Four  = 3 + 1 => 4
	Four
	Six = iota+1
)*/

const (
	// iota: 0, One: 1 (type: int64)
	One int64 = iota + 1
	// iota: 1, Two: 2 (type: int64)
	// Two will be declared as if:
	// Two int64 = iota + 1
	Two
	// iota: 2, Four: 4 (type: int32)
	Four int32 = iota + 2
	// iota: 3, Five: 5 (type: int32)
	// Five will be declared as if:
	// Five int32 = iota + 2
	Five
	// (type: int)
	Six = 6
	// (type: int)
	// Seven will be declared as if:
	// Seven = 6
	Seven
)

type Month int

// Bitwise operations
const (
	// 1 << 0 ==> 1
	January  Month = 1 << iota
	February       // 1 << 1 ==> 2
	March          // 1 << 2 ==> 4
	April          // 1 << 3 ==> 8
	May            // 1 << 4 ==> 16
	June           // ...
	July
	August
	September
	October
	November
	December
	// Break the iota chain here.
	// AllMonths will have only
	// the assigned month values,
	// not the iota's.
	AllMonths = January | February |
		March | April | May | June |
		July | August | September |
		October | November |
		December
)

/*const (
	// Active = 0, Running = 100
	Active, Running = iota, iota + 100
	// Passive = 1, Stopped = 101
	Passive, Stopped
	// You can't declare like this.
	// The last expression will be
	// repeated
	// CantDeclare
	// But, you can reset
	// the last expression
	Reset = iota
	// You can use any other
	// expression even without iota
	AnyOther = 10
)*/

// type Activity int
//const (
//    Sleeping = iota
//    Walking
//    Running
//)
//func main() {
//    var activity Activity
//    // activity initialized to
//    // its zero-value of int
//    // which is Sleeping
//}

// 0 is a zero-value for integers. So, you can’t know whether the Activity is initialized or not; Is it really in the Sleeping state?

// Use “iota + 1” to be sure that the enum type is initialized.
type Activity int

// iota + 1 trick
const (
	Sleeping = iota + 1
	Walking
	Reset = iota
	Running
)

func (day Weekday) String() string {
	names := [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday"}
	// → `day`: It's one of the
	// values of Weekday constants.
	// If the constant is Sunday,
	// then day is 0.
	//
	// prevent panicking in case of
	// `day` is out of range of Weekday
	if day < Sunday || day > Saturday {
		return "Unknown"
	}
	// return the name of a Weekday
	// constant from the names array
	// above.
	return names[day]
}

func (day Weekday) Weekend() bool {
	if day == Saturday || day == Sunday {
		return true
	}
	return false
}

func Iota() {
	fmt.Println(One, Two, Four, Six, Five, Six, Seven)
	var activity Activity
	// activity will be zero,
	// so it's not initialized
	activity = Sleeping
	// now you know that it's been
	// initialized
	fmt.Println(activity)
}
