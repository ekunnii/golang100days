package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2} //has type Vertex
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p5  = &Vertex{1, 2}
)

func main() {
    //Page 6
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1])
    fmt.Println(a)
    
    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)
    
    //Page 5
    fmt.Println(v1, p5, v2, v3)
    
	//Page 3
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	//Page 4
	p2 := &v
	p2.X = 1e9
	fmt.Println(v)

	//Page 2
	fmt.Println(Vertex{1, 2})

	//Page 1
	i, j := 42, 2701

	p := &i         //point to i
	fmt.Println(*p) //read i through the pointer
	*p = 21         //set i through the pointer
	fmt.Println(i)  //see the new value of i

	p = &j         //point to j
	*p = *p / 37   //divide j through the pointer
	fmt.Println(j) // see the new value of j

}
