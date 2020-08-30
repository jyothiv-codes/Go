package main

import "fmt"

type Vertex struct  {
	X int
	Y int
}

func main() {
	var v1=Vertex{2,5}
	var v2=Vertex{12,9}
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v1,v2)
	v1.Y=6
	fmt.Println(v1)
	fmt.Println(v1.X)
	v3:=Vertex{4,3}
	fmt.Println(v3)
}
