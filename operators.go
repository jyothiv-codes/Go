/*
1. int can't work here as value for if condition unlike few other langs
2. else is not compulsory
3. else if 
*/
package main
import "fmt"
func main(){
	var a int 
	a=3 
	if a==3 {
		fmt.Println("True")
	}else{
		fmt.Println("False")
	}
	if !true {
		fmt.Println("Negated")
	}
	var b float64
	b=4
	fmt.Println(b)
	if b==4 {
		fmt.Println(b)
	}
	i:=7
	for i<20{
		fmt.Println(i)
		i++
	}
	var num int
	fmt.Println("Enter a number: ")
	fmt.Scanln(&num)
	if num%5==0 && num%3==0{
		fmt.Println("Divisible by 5 and 3")
	}else if num%5==0{
		fmt.Println("Divisible by 5")
	} else if{
		fmt.Println("Divisible by 3")
	}else{
		fmt.Println("Not divisible by both 5 and 3")
	}

}


