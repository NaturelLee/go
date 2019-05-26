package numbTest

import (
	"fmt"
	"strconv"
)

func ConvNumbers() {
	fmt.Println("numbTest package is in mygo/")
	str := strconv.Itoa(int(1))
	fmt.Println("int => string", str)

	str2 := strconv.FormatInt(int64(1), 10)
	fmt.Println("int64 => string", str2)

	str3 := strconv.FormatFloat(float64(0.8), 'f', -1, 32)
	fmt.Println("float32 => string", str3)

	i, _ := strconv.Atoi("10")
	fmt.Println("string => int", i)

	i64, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println("string => int64", i64)

	fl32, _ := strconv.ParseFloat("3.1415926", 32/64)
	fmt.Println("string => float64", fl32)

}
