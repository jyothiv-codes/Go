/*
learning to use error messages
will also learn about := and = a little 

*/

package main
import (
	"fmt";
	"math";
	"errors"
)
	
func sqRoot(num float64) (float64,error){
	if num<0{
		return 0,errors.New("Negative number sent!!")
	}else{
		return math.Sqrt(num),nil
	}
}
func main(){
	var num float64
	fmt.Println("Enter a number: ")
	fmt.Scanln(&num)
	result, err:=sqRoot(num)
	if err== nil{
		fmt.Println(result)
	}else{
		fmt.Println(err)
	}
}