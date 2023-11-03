package main

import "fmt"

type ValueChecker interface {
	IsInt() (bool, int)
	IsMoreThanOrEqualOne() bool
	IsLessThanOrEqualTen() bool
	Check() bool
}

type IntervalValueChecker struct {
	value any
}

func (ivc *IntervalValueChecker) IsInt() (bool, int) {
	valueInt, ok := ivc.value.(int)
	if !ok {
		return false, 0
	}

	return true, valueInt
}

func (ivc *IntervalValueChecker) IsMoreThanOrEqualOne() bool {
	return ivc.value.(int) >= 1
}

func (ivc *IntervalValueChecker) IsLessThanOrEqualTen() bool {
	return ivc.value.(int) <= 10
}

func (ivc *IntervalValueChecker) Check() bool {
	res, value := ivc.IsInt()
	if !res {
		return res
	}

	ivc.value = value
	return ivc.IsMoreThanOrEqualOne() && ivc.IsLessThanOrEqualTen()
}

func main() {
	value1 := IntervalValueChecker{value: 5}
	value2 := IntervalValueChecker{value: "51"}

	fmt.Println("Valid value check:", value1.Check())
	fmt.Println("Invalid value check:", value2.Check())
}
