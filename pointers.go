/*
have to work on dereferencing a pointer
*/
package main
import "fmt"
func main(){
	//part 1
	var num int 
	fmt.Println(&num)
	fmt.Println("%x",&num)
	fmt.Printf("Address of a variable: %x\n", &num  )

	//part 2
	var pointer *int 
	var a=25;
	fmt.Println("\n\nPart 2")
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(pointer)
	fmt.Println(*pointer)
	pointer=&a 
	fmt.Println(pointer)
	fmt.Println(*pointer)
	
}
