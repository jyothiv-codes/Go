
/*
1. for loop - general
2. for loop - pattern */

package main
import "fmt"
func main(){
	sum:=0 
	for i:=0;i<10;i++{
		fmt.Println(i)
		sum+=1
	}
	fmt.Println("Sum :",sum)
	for i:=0;i<5;i++{
		for j:=i;j<5;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
	var a int
	var b int 
	var sum int 
	fmt.Println("enter the value of a: ")
	fmt.Scanln(&a)
	fmt.Println("enter the value of b: ")
	fmt.Scanln(&b)

}