/*
1. usage of fmt, strings and reflect
2. user defined function: return default values, computed values
3. check if chars are present in string
4. length of string


*/

package main 
import (
	"fmt" 
	"strings"
	"reflect")
func SimpleFunction() {
	fmt.Println("Hello World")
}
func amount() int{
	return 15
}
func total(x int,y int) int{
	return x+y
}
func str(st string) string{
	return st
}
func main(){
	SimpleFunction()
	fmt.Println(amount())
	fmt.Println(total(2,3))
	//fmt.Println(str("hello"))
	str1 := "Everyday offers a new beginning...."
    res1 := strings.ContainsAny(str1, "abcde") 
    fmt.Println(res1)
    fmt.Println(len(str1))
    fmt.Println(reflect.TypeOf(str1))
}