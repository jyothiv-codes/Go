/*
1. break statement
2. a++
3. boolean values
4. Default value of a variable
*/

package main
import "fmt"
func main(){
	var value int
	fmt.Println("Enter a value: ")
	fmt.Scanln(&value)
	if value%2==0{
		fmt.Println("Even number")
	}else{
		fmt.Println("Odd number")
	}
	var x int 
	fmt.Println("Default value of a variable: ",x)
	x++
	fmt.Println(x)
	if true{
		fmt.Println("True value")
	}
	if !false{
		fmt.Println("!false")
	}
	var i int
	for i=0;i<10;i++{
		if i==4{
			break
		}else{
			fmt.Println(i)
		}
	}
}