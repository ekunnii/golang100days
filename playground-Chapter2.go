package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func Sqrt(x float64) float64 {
	//Newton's method
	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
    //Page 13
    fmt.Println("counting")
    
    for i:=0; i < 10; i++ {
        defer fmt.Println(i)
    }
    
    fmt.Println("done")
    
    //Page 12
    defer fmt.Println("world")
    
    fmt.Println("Hello")
    
	//Page 11
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	//Page 10
	fmt.Println("When's Sturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}

	//Page 9
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		//freebsd, openbsd,
		//Plan9, windows
		fmt.Printf("%s.\n", os)
	}

	//Page 8
	fmt.Println(Sqrt(2))
	fmt.Println(sqrt(2))
	//page 6 & 7
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	//page 5
	fmt.Println(sqrt(2), sqrt(-4))

	//page 4
	//    for {
	//    }

	//page 3
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println(sum3)

	//page 2
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	//page 1
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

}
