package semantics

import (
	"fmt"
	"math"
)

// Go错误优雅处理

type AreaError struct {
	err    string
	radius float64
	length int
	width  int
}

func (e *AreaError) Error() string {
	return fmt.Sprintf("radius:%0.2f:%s", e.radius, e.err)
}

func (e *AreaError) lengthNegative() bool {
	return e.length < 0
}

func (e *AreaError) widthNegative() bool {
	return e.width < 0
}

func CircleArea(radius float64) (float64, error) {
	if radius < 0 {
		// 明显自定义的错误 更加灵活表达错误信息
		// errorf := fmt.Errorf("radius %v is negative ", radius)
		return 0, &AreaError{"radius is negative", radius, nil, nil}
	}
	return math.Pi * radius * radius, nil
}

func Err() {
	radius := -20.0
	area, err := CircleArea(radius)
	if err != nil {
		if err, ok := err.(*AreaError); ok {
			fmt.Printf("Radius %0.2f is less than zero", err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of rectangle1 %0.2f", area)
}
