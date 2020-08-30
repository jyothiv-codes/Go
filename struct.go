/*
Struct 

1. initialisation- different ways
2. modifying values
3. Understand how null structs get printed
4. Since fmt.Println is assigned a variable within main, we can use both the variable and the regular function
*/

package main
import "fmt"

type Vertex struct  {
	X int
	Y int
}
type diagrams struct {
	shape string
	colour string 
	quantity int 
}

func main() {
	var print=fmt.Println
	var v1=Vertex{2,5}
	var v2=Vertex{12,9}
	fmt.Println(v1)
	fmt.Println(v1,v2)
	v1.Y=6
	fmt.Println(v1)
	fmt.Println(v1.Y)
	v3:=Vertex{4,3}
	fmt.Println(v3)
	v4:=new(Vertex)
	fmt.Println("V4",v4)
	fmt.Println(v4.Y)
	print("Here we are checking the assignment of fmt.Println to a variable")
	fmt.Println("Here is regular fmt.Println")
	print("Next: assigning structs to one another and then modifying the values: ")
	d1:=diagrams{"triangle","green",3}
	d2:=d1 
	d2.shape="Circle"
	print(d1,d2)// we understand that d2 objects and d1 objects differ completely (in terms of memory)
	

}
